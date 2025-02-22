# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: swim.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'swim.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nswim.proto\x12\x04swim\"3\n\x0bPingRequest\x12\x11\n\tsender_id\x18\x01 \x01(\t\x12\x11\n\ttarget_id\x18\x02 \x01(\t\" \n\x0cPingResponse\x12\x10\n\x08is_alive\x18\x01 \x01(\x08\"S\n\x13IndirectPingRequest\x12\x14\n\x0crequester_id\x18\x01 \x01(\t\x12\x11\n\ttarget_id\x18\x02 \x01(\t\x12\x13\n\x0bproxy_nodes\x18\x03 \x03(\t\"\'\n\x14IndirectPingResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"-\n\x13\x46\x61ilureNotification\x12\x16\n\x0e\x66\x61iled_node_id\x18\x01 \x01(\t2\x87\x01\n\x0f\x46\x61ilureDetector\x12-\n\x04Ping\x12\x11.swim.PingRequest\x1a\x12.swim.PingResponse\x12\x45\n\x0cIndirectPing\x12\x19.swim.IndirectPingRequest\x1a\x1a.swim.IndirectPingResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'swim_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_PINGREQUEST']._serialized_start=20
  _globals['_PINGREQUEST']._serialized_end=71
  _globals['_PINGRESPONSE']._serialized_start=73
  _globals['_PINGRESPONSE']._serialized_end=105
  _globals['_INDIRECTPINGREQUEST']._serialized_start=107
  _globals['_INDIRECTPINGREQUEST']._serialized_end=190
  _globals['_INDIRECTPINGRESPONSE']._serialized_start=192
  _globals['_INDIRECTPINGRESPONSE']._serialized_end=231
  _globals['_FAILURENOTIFICATION']._serialized_start=233
  _globals['_FAILURENOTIFICATION']._serialized_end=278
  _globals['_FAILUREDETECTOR']._serialized_start=281
  _globals['_FAILUREDETECTOR']._serialized_end=416
# @@protoc_insertion_point(module_scope)
