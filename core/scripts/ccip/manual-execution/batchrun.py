#!/usr/bin/env python3

import subprocess
import os
import tempfile
import json

# go build -o ccipmanuaalexec
# ./batchrun.py

binary_name = "./ccipmanualexec"
input_msgs = "msgs.csv"


# global config
src_rpc = os.environ["src_rpc"]
dest_rpc = os.environ["dest_rpc"]
dest_owner_key = os.environ["dest_owner_key"]
commit_store = os.environ["commit_store"]
off_ramp = os.environ["off_ramp"]
dest_start_block = os.environ["dest_start_block"]
source_start_block = os.environ["source_start_block"]
dest_deployed_at = os.environ["dest_deployed_at"]
gas_limit_override = os.environ["gas_limit_override"]

msgs = open(input_msgs, "r").read().split("\n")[1:]
for i, msg in enumerate(msgs):
    # per msg config
    parts = msg.split(",")
    if len(parts) != 2:
        continue

    ccip_send_tx = parts[1]
    msg_id = parts[0]

    print("[%d/%d] >>> %s %s" % (i, len(msgs), ccip_send_tx, msg_id))

    config = {
        "ccip_send_tx": ccip_send_tx,
        "msg_id": msg_id,
        "src_rpc": src_rpc,
        "dest_rpc": dest_rpc,
        "dest_owner_key": dest_owner_key,
        "commit_store": commit_store,
        "off_ramp": off_ramp,
        "dest_start_block": dest_start_block,
        "source_start_block": source_start_block,
        "dest_deployed_at": dest_deployed_at,
        "gas_limit_override": gas_limit_override
    }
    json_config = json.dumps(config)

    json_config_file = tempfile.NamedTemporaryFile()
    with open(json_config_file.name, 'w') as f:
        f.write(json_config)

    try:
        subprocess.run([binary_name, json_config_file.name])
    except subprocess.CalledProcessError as e:
        print("called process error: ", err)

