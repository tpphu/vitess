#!/usr/bin/env python

# Copyright 2015 Google Inc. All rights reserved.
# Use of this source code is governed by a BSD-style license that can
# be found in the LICENSE file.

"""This sample test demonstrates how to setup and teardown a test
database with the associated Vitess processes. It is meant to be used
as a template for unit tests, by developers writing applications
on top of Vitess. The recommended workflow is to have the schema for the
database checked into a source control system, along with the application
layer code. Then unit tests can launch a test cluster for unit tests
(using local_database.py) pointing at that schema, and run all their unit
tests.

This unit test is written in python, but we don't depend on this at all.
We just execute py/vttest/local_database.py, as we would in any language.
Then we wait for the JSON config string, and we know the address of the
vtgate process. At the end of the test, sending a line to the underlying
process makes it clean up.
"""

import json
import os
import subprocess
import urllib

import unittest

from vttest import fakezk_config

import utils
import environment


class TestMysqlctl(unittest.TestCase):

  def test_fakezk_config(self):
    """This unit test makes sure the fake zk config is correct."""

    c = fakezk_config.FakeZkConfig(mysql_port=3307, cell='test_cell')
    c.add_shard('video', '0', 6711)
    c.add_shard('user', '-80', 6712)
    c.add_shard('user', '80-', 6713)
    c.add_redirect('ingestion', 'ingested')
    json_data = c.as_json()
    parsed_data = json.loads(json_data)
    self.assertEqual(
        '-80',
        parsed_data['/zk/test_cell/vt/ns/user']
        ['partitions'][0]['shard_references'][0]['name'])
    self.assertEqual(
        'gA==',  # base 64 of 0x80
        parsed_data['/zk/test_cell/vt/ns/user']
        ['partitions'][0]['shard_references'][1]['key_range']
        ['start'])
    self.assertNotIn(
        'end',
        parsed_data['/zk/test_cell/vt/ns/user']
        ['partitions'][0]['shard_references'][1]['key_range'])
    self.assertEqual(
        6711,
        parsed_data['/zk/test_cell/vt/ns/video/0/master']['entries'][0]
        ['port_map']['vt'])
    self.assertEqual(
        'ingested',
        parsed_data['/zk/test_cell/vt/ns/ingestion']['served_from'][0]
        ['keyspace'])

  def test_standalone_vtgate(self):
    """Sample test for run_local_database.py(vtgate) as a standalone process."""
    self._test_standalone(use_vtcombo=False)

  def test_standalone_vtcombo(self):
    """Sample test for run_local_database.py(vtcombo) as a standalone process."""
    self._test_standalone(use_vtcombo=True)

  def _test_standalone(self, use_vtcombo):
    """Sample test for run_local_database.py as a standalone process."""

    # launch a backend database based on the provided topology and schema
    port = environment.reserve_ports(1)
    args = [environment.run_local_database,
            '--port', str(port),
            '--topology',
            'test_keyspace/-80:test_keyspace_0,'
            'test_keyspace/80-:test_keyspace_1',
            '--schema_dir', os.path.join(environment.vttop, 'test',
                                         'vttest_schema'),
    ]
    if use_vtcombo:
      args.append('--use_vtcombo')
    sp = subprocess.Popen(args, stdin=subprocess.PIPE, stdout=subprocess.PIPE)
    config = json.loads(sp.stdout.readline())

    # gather the vars for the vtgate process
    url = 'http://localhost:%d/debug/vars' % config['port']
    f = urllib.urlopen(url)
    data = f.read()
    f.close()
    json_vars = json.loads(data)
    process_name = 'vtgate'
    if use_vtcombo:
      process_name = 'vtcombo'
    self.assertIn(process_name, json_vars['cmdline'][0])

    # to test vtcombo:
    # ./vttest_sample_test.py -v -d
    # go install && vtcombo -port 15010 -grpc_port 15011 -service_map grpc-vtgateservice -topology test_keyspace/-80:test_keyspace_0,test_keyspace/80-:test_keyspace_1 -mycnf_server_id 1 -mycnf_socket_file $VTDATAROOT/vttest*/vt_0000000001/mysql.sock -db-config-dba-uname vt_dba -db-config-dba-charset utf8 -db-config-app-uname vt_app -db-config-app-charset utf8 -alsologtostderr
    # vtctl -vtgate_protocol grpc VtGateExecuteShard -server localhost:15011 -keyspace test_keyspace -shards -80 -tablet_type master "select 1 from dual"
    if use_vtcombo:
      utils.pause('good time to test vtcombo with database running')
    else:
      utils.pause('good time to test vtgate with database running')

    # and we're done, clean-up
    sp.stdin.write('\n')
    sp.wait()

if __name__ == '__main__':
  utils.main()