# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mcp/v1alpha1/mcp.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from mcp.v1alpha1 import resource_pb2 as mcp_dot_v1alpha1_dot_resource__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='mcp/v1alpha1/mcp.proto',
  package='istio.mcp.v1alpha1',
  syntax='proto3',
  serialized_options=_b('Z\031istio.io/api/mcp/v1alpha1\250\342\036\001'),
  serialized_pb=_b('\n\x16mcp/v1alpha1/mcp.proto\x12\x12istio.mcp.v1alpha1\x1a\x17google/rpc/status.proto\x1a\x14gogoproto/gogo.proto\x1a\x1bmcp/v1alpha1/resource.proto\"\x8e\x01\n\x08SinkNode\x12\n\n\x02id\x18\x01 \x01(\t\x12\x42\n\x0b\x61nnotations\x18\x02 \x03(\x0b\x32-.istio.mcp.v1alpha1.SinkNode.AnnotationsEntry\x1a\x32\n\x10\x41nnotationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xae\x01\n\x11MeshConfigRequest\x12\x14\n\x0cversion_info\x18\x01 \x01(\t\x12/\n\tsink_node\x18\x02 \x01(\x0b\x32\x1c.istio.mcp.v1alpha1.SinkNode\x12\x10\n\x08type_url\x18\x03 \x01(\t\x12\x16\n\x0eresponse_nonce\x18\x04 \x01(\t\x12(\n\x0c\x65rror_detail\x18\x05 \x01(\x0b\x32\x12.google.rpc.Status\"\x82\x01\n\x12MeshConfigResponse\x12\x14\n\x0cversion_info\x18\x01 \x01(\t\x12\x35\n\tresources\x18\x02 \x03(\x0b\x32\x1c.istio.mcp.v1alpha1.ResourceB\x04\xc8\xde\x1f\x00\x12\x10\n\x08type_url\x18\x03 \x01(\t\x12\r\n\x05nonce\x18\x04 \x01(\t\"\xd5\x02\n\x1cIncrementalMeshConfigRequest\x12/\n\tsink_node\x18\x01 \x01(\x0b\x32\x1c.istio.mcp.v1alpha1.SinkNode\x12\x10\n\x08type_url\x18\x02 \x01(\t\x12p\n\x19initial_resource_versions\x18\x03 \x03(\x0b\x32M.istio.mcp.v1alpha1.IncrementalMeshConfigRequest.InitialResourceVersionsEntry\x12\x16\n\x0eresponse_nonce\x18\x04 \x01(\t\x12(\n\x0c\x65rror_detail\x18\x05 \x01(\x0b\x32\x12.google.rpc.Status\x1a>\n\x1cInitialResourceVersionsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x9d\x01\n\x1dIncrementalMeshConfigResponse\x12\x1b\n\x13system_version_info\x18\x01 \x01(\t\x12\x35\n\tresources\x18\x02 \x03(\x0b\x32\x1c.istio.mcp.v1alpha1.ResourceB\x04\xc8\xde\x1f\x00\x12\x19\n\x11removed_resources\x18\x03 \x03(\t\x12\r\n\x05nonce\x18\x04 \x01(\t\"\xd4\x02\n\x10RequestResources\x12/\n\tsink_node\x18\x01 \x01(\x0b\x32\x1c.istio.mcp.v1alpha1.SinkNode\x12\x12\n\ncollection\x18\x02 \x01(\t\x12\x64\n\x19initial_resource_versions\x18\x03 \x03(\x0b\x32\x41.istio.mcp.v1alpha1.RequestResources.InitialResourceVersionsEntry\x12\x16\n\x0eresponse_nonce\x18\x04 \x01(\t\x12(\n\x0c\x65rror_detail\x18\x05 \x01(\x0b\x32\x12.google.rpc.Status\x12\x13\n\x0bincremental\x18\x06 \x01(\x08\x1a>\n\x1cInitialResourceVersionsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xb2\x01\n\tResources\x12\x1b\n\x13system_version_info\x18\x01 \x01(\t\x12\x12\n\ncollection\x18\x02 \x01(\t\x12\x35\n\tresources\x18\x03 \x03(\x0b\x32\x1c.istio.mcp.v1alpha1.ResourceB\x04\xc8\xde\x1f\x00\x12\x19\n\x11removed_resources\x18\x04 \x03(\t\x12\r\n\x05nonce\x18\x05 \x01(\t\x12\x13\n\x0bincremental\x18\x06 \x01(\x08\x32\x9d\x02\n\x1b\x41ggregatedMeshConfigService\x12p\n\x19StreamAggregatedResources\x12%.istio.mcp.v1alpha1.MeshConfigRequest\x1a&.istio.mcp.v1alpha1.MeshConfigResponse\"\x00(\x01\x30\x01\x12\x8b\x01\n\x1eIncrementalAggregatedResources\x12\x30.istio.mcp.v1alpha1.IncrementalMeshConfigRequest\x1a\x31.istio.mcp.v1alpha1.IncrementalMeshConfigResponse\"\x00(\x01\x30\x01\x32v\n\x0eResourceSource\x12\x64\n\x17\x45stablishResourceStream\x12$.istio.mcp.v1alpha1.RequestResources\x1a\x1d.istio.mcp.v1alpha1.Resources\"\x00(\x01\x30\x01\x32t\n\x0cResourceSink\x12\x64\n\x17\x45stablishResourceStream\x12\x1d.istio.mcp.v1alpha1.Resources\x1a$.istio.mcp.v1alpha1.RequestResources\"\x00(\x01\x30\x01\x42\x1fZ\x19istio.io/api/mcp/v1alpha1\xa8\xe2\x1e\x01\x62\x06proto3')
  ,
  dependencies=[google_dot_rpc_dot_status__pb2.DESCRIPTOR,gogoproto_dot_gogo__pb2.DESCRIPTOR,mcp_dot_v1alpha1_dot_resource__pb2.DESCRIPTOR,])




_SINKNODE_ANNOTATIONSENTRY = _descriptor.Descriptor(
  name='AnnotationsEntry',
  full_name='istio.mcp.v1alpha1.SinkNode.AnnotationsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mcp.v1alpha1.SinkNode.AnnotationsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mcp.v1alpha1.SinkNode.AnnotationsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=215,
  serialized_end=265,
)

_SINKNODE = _descriptor.Descriptor(
  name='SinkNode',
  full_name='istio.mcp.v1alpha1.SinkNode',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='istio.mcp.v1alpha1.SinkNode.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='annotations', full_name='istio.mcp.v1alpha1.SinkNode.annotations', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SINKNODE_ANNOTATIONSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=123,
  serialized_end=265,
)


_MESHCONFIGREQUEST = _descriptor.Descriptor(
  name='MeshConfigRequest',
  full_name='istio.mcp.v1alpha1.MeshConfigRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version_info', full_name='istio.mcp.v1alpha1.MeshConfigRequest.version_info', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sink_node', full_name='istio.mcp.v1alpha1.MeshConfigRequest.sink_node', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type_url', full_name='istio.mcp.v1alpha1.MeshConfigRequest.type_url', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='response_nonce', full_name='istio.mcp.v1alpha1.MeshConfigRequest.response_nonce', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error_detail', full_name='istio.mcp.v1alpha1.MeshConfigRequest.error_detail', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=268,
  serialized_end=442,
)


_MESHCONFIGRESPONSE = _descriptor.Descriptor(
  name='MeshConfigResponse',
  full_name='istio.mcp.v1alpha1.MeshConfigResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version_info', full_name='istio.mcp.v1alpha1.MeshConfigResponse.version_info', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='istio.mcp.v1alpha1.MeshConfigResponse.resources', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type_url', full_name='istio.mcp.v1alpha1.MeshConfigResponse.type_url', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='nonce', full_name='istio.mcp.v1alpha1.MeshConfigResponse.nonce', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=445,
  serialized_end=575,
)


_INCREMENTALMESHCONFIGREQUEST_INITIALRESOURCEVERSIONSENTRY = _descriptor.Descriptor(
  name='InitialResourceVersionsEntry',
  full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.InitialResourceVersionsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.InitialResourceVersionsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.InitialResourceVersionsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=857,
  serialized_end=919,
)

_INCREMENTALMESHCONFIGREQUEST = _descriptor.Descriptor(
  name='IncrementalMeshConfigRequest',
  full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sink_node', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.sink_node', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type_url', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.type_url', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='initial_resource_versions', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.initial_resource_versions', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='response_nonce', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.response_nonce', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error_detail', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigRequest.error_detail', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_INCREMENTALMESHCONFIGREQUEST_INITIALRESOURCEVERSIONSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=578,
  serialized_end=919,
)


_INCREMENTALMESHCONFIGRESPONSE = _descriptor.Descriptor(
  name='IncrementalMeshConfigResponse',
  full_name='istio.mcp.v1alpha1.IncrementalMeshConfigResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='system_version_info', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigResponse.system_version_info', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigResponse.resources', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='removed_resources', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigResponse.removed_resources', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='nonce', full_name='istio.mcp.v1alpha1.IncrementalMeshConfigResponse.nonce', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=922,
  serialized_end=1079,
)


_REQUESTRESOURCES_INITIALRESOURCEVERSIONSENTRY = _descriptor.Descriptor(
  name='InitialResourceVersionsEntry',
  full_name='istio.mcp.v1alpha1.RequestResources.InitialResourceVersionsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mcp.v1alpha1.RequestResources.InitialResourceVersionsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mcp.v1alpha1.RequestResources.InitialResourceVersionsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=857,
  serialized_end=919,
)

_REQUESTRESOURCES = _descriptor.Descriptor(
  name='RequestResources',
  full_name='istio.mcp.v1alpha1.RequestResources',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sink_node', full_name='istio.mcp.v1alpha1.RequestResources.sink_node', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collection', full_name='istio.mcp.v1alpha1.RequestResources.collection', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='initial_resource_versions', full_name='istio.mcp.v1alpha1.RequestResources.initial_resource_versions', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='response_nonce', full_name='istio.mcp.v1alpha1.RequestResources.response_nonce', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error_detail', full_name='istio.mcp.v1alpha1.RequestResources.error_detail', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='incremental', full_name='istio.mcp.v1alpha1.RequestResources.incremental', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_REQUESTRESOURCES_INITIALRESOURCEVERSIONSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1082,
  serialized_end=1422,
)


_RESOURCES = _descriptor.Descriptor(
  name='Resources',
  full_name='istio.mcp.v1alpha1.Resources',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='system_version_info', full_name='istio.mcp.v1alpha1.Resources.system_version_info', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collection', full_name='istio.mcp.v1alpha1.Resources.collection', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='istio.mcp.v1alpha1.Resources.resources', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='removed_resources', full_name='istio.mcp.v1alpha1.Resources.removed_resources', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='nonce', full_name='istio.mcp.v1alpha1.Resources.nonce', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='incremental', full_name='istio.mcp.v1alpha1.Resources.incremental', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1425,
  serialized_end=1603,
)

_SINKNODE_ANNOTATIONSENTRY.containing_type = _SINKNODE
_SINKNODE.fields_by_name['annotations'].message_type = _SINKNODE_ANNOTATIONSENTRY
_MESHCONFIGREQUEST.fields_by_name['sink_node'].message_type = _SINKNODE
_MESHCONFIGREQUEST.fields_by_name['error_detail'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_MESHCONFIGRESPONSE.fields_by_name['resources'].message_type = mcp_dot_v1alpha1_dot_resource__pb2._RESOURCE
_INCREMENTALMESHCONFIGREQUEST_INITIALRESOURCEVERSIONSENTRY.containing_type = _INCREMENTALMESHCONFIGREQUEST
_INCREMENTALMESHCONFIGREQUEST.fields_by_name['sink_node'].message_type = _SINKNODE
_INCREMENTALMESHCONFIGREQUEST.fields_by_name['initial_resource_versions'].message_type = _INCREMENTALMESHCONFIGREQUEST_INITIALRESOURCEVERSIONSENTRY
_INCREMENTALMESHCONFIGREQUEST.fields_by_name['error_detail'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_INCREMENTALMESHCONFIGRESPONSE.fields_by_name['resources'].message_type = mcp_dot_v1alpha1_dot_resource__pb2._RESOURCE
_REQUESTRESOURCES_INITIALRESOURCEVERSIONSENTRY.containing_type = _REQUESTRESOURCES
_REQUESTRESOURCES.fields_by_name['sink_node'].message_type = _SINKNODE
_REQUESTRESOURCES.fields_by_name['initial_resource_versions'].message_type = _REQUESTRESOURCES_INITIALRESOURCEVERSIONSENTRY
_REQUESTRESOURCES.fields_by_name['error_detail'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_RESOURCES.fields_by_name['resources'].message_type = mcp_dot_v1alpha1_dot_resource__pb2._RESOURCE
DESCRIPTOR.message_types_by_name['SinkNode'] = _SINKNODE
DESCRIPTOR.message_types_by_name['MeshConfigRequest'] = _MESHCONFIGREQUEST
DESCRIPTOR.message_types_by_name['MeshConfigResponse'] = _MESHCONFIGRESPONSE
DESCRIPTOR.message_types_by_name['IncrementalMeshConfigRequest'] = _INCREMENTALMESHCONFIGREQUEST
DESCRIPTOR.message_types_by_name['IncrementalMeshConfigResponse'] = _INCREMENTALMESHCONFIGRESPONSE
DESCRIPTOR.message_types_by_name['RequestResources'] = _REQUESTRESOURCES
DESCRIPTOR.message_types_by_name['Resources'] = _RESOURCES
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SinkNode = _reflection.GeneratedProtocolMessageType('SinkNode', (_message.Message,), {

  'AnnotationsEntry' : _reflection.GeneratedProtocolMessageType('AnnotationsEntry', (_message.Message,), {
    'DESCRIPTOR' : _SINKNODE_ANNOTATIONSENTRY,
    '__module__' : 'mcp.v1alpha1.mcp_pb2'
    # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.SinkNode.AnnotationsEntry)
    })
  ,
  'DESCRIPTOR' : _SINKNODE,
  '__module__' : 'mcp.v1alpha1.mcp_pb2'
  # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.SinkNode)
  })
_sym_db.RegisterMessage(SinkNode)
_sym_db.RegisterMessage(SinkNode.AnnotationsEntry)

MeshConfigRequest = _reflection.GeneratedProtocolMessageType('MeshConfigRequest', (_message.Message,), {
  'DESCRIPTOR' : _MESHCONFIGREQUEST,
  '__module__' : 'mcp.v1alpha1.mcp_pb2'
  # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.MeshConfigRequest)
  })
_sym_db.RegisterMessage(MeshConfigRequest)

MeshConfigResponse = _reflection.GeneratedProtocolMessageType('MeshConfigResponse', (_message.Message,), {
  'DESCRIPTOR' : _MESHCONFIGRESPONSE,
  '__module__' : 'mcp.v1alpha1.mcp_pb2'
  # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.MeshConfigResponse)
  })
_sym_db.RegisterMessage(MeshConfigResponse)

IncrementalMeshConfigRequest = _reflection.GeneratedProtocolMessageType('IncrementalMeshConfigRequest', (_message.Message,), {

  'InitialResourceVersionsEntry' : _reflection.GeneratedProtocolMessageType('InitialResourceVersionsEntry', (_message.Message,), {
    'DESCRIPTOR' : _INCREMENTALMESHCONFIGREQUEST_INITIALRESOURCEVERSIONSENTRY,
    '__module__' : 'mcp.v1alpha1.mcp_pb2'
    # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.IncrementalMeshConfigRequest.InitialResourceVersionsEntry)
    })
  ,
  'DESCRIPTOR' : _INCREMENTALMESHCONFIGREQUEST,
  '__module__' : 'mcp.v1alpha1.mcp_pb2'
  # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.IncrementalMeshConfigRequest)
  })
_sym_db.RegisterMessage(IncrementalMeshConfigRequest)
_sym_db.RegisterMessage(IncrementalMeshConfigRequest.InitialResourceVersionsEntry)

IncrementalMeshConfigResponse = _reflection.GeneratedProtocolMessageType('IncrementalMeshConfigResponse', (_message.Message,), {
  'DESCRIPTOR' : _INCREMENTALMESHCONFIGRESPONSE,
  '__module__' : 'mcp.v1alpha1.mcp_pb2'
  # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.IncrementalMeshConfigResponse)
  })
_sym_db.RegisterMessage(IncrementalMeshConfigResponse)

RequestResources = _reflection.GeneratedProtocolMessageType('RequestResources', (_message.Message,), {

  'InitialResourceVersionsEntry' : _reflection.GeneratedProtocolMessageType('InitialResourceVersionsEntry', (_message.Message,), {
    'DESCRIPTOR' : _REQUESTRESOURCES_INITIALRESOURCEVERSIONSENTRY,
    '__module__' : 'mcp.v1alpha1.mcp_pb2'
    # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.RequestResources.InitialResourceVersionsEntry)
    })
  ,
  'DESCRIPTOR' : _REQUESTRESOURCES,
  '__module__' : 'mcp.v1alpha1.mcp_pb2'
  # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.RequestResources)
  })
_sym_db.RegisterMessage(RequestResources)
_sym_db.RegisterMessage(RequestResources.InitialResourceVersionsEntry)

Resources = _reflection.GeneratedProtocolMessageType('Resources', (_message.Message,), {
  'DESCRIPTOR' : _RESOURCES,
  '__module__' : 'mcp.v1alpha1.mcp_pb2'
  # @@protoc_insertion_point(class_scope:istio.mcp.v1alpha1.Resources)
  })
_sym_db.RegisterMessage(Resources)


DESCRIPTOR._options = None
_SINKNODE_ANNOTATIONSENTRY._options = None
_MESHCONFIGRESPONSE.fields_by_name['resources']._options = None
_INCREMENTALMESHCONFIGREQUEST_INITIALRESOURCEVERSIONSENTRY._options = None
_INCREMENTALMESHCONFIGRESPONSE.fields_by_name['resources']._options = None
_REQUESTRESOURCES_INITIALRESOURCEVERSIONSENTRY._options = None
_RESOURCES.fields_by_name['resources']._options = None

_AGGREGATEDMESHCONFIGSERVICE = _descriptor.ServiceDescriptor(
  name='AggregatedMeshConfigService',
  full_name='istio.mcp.v1alpha1.AggregatedMeshConfigService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1606,
  serialized_end=1891,
  methods=[
  _descriptor.MethodDescriptor(
    name='StreamAggregatedResources',
    full_name='istio.mcp.v1alpha1.AggregatedMeshConfigService.StreamAggregatedResources',
    index=0,
    containing_service=None,
    input_type=_MESHCONFIGREQUEST,
    output_type=_MESHCONFIGRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='IncrementalAggregatedResources',
    full_name='istio.mcp.v1alpha1.AggregatedMeshConfigService.IncrementalAggregatedResources',
    index=1,
    containing_service=None,
    input_type=_INCREMENTALMESHCONFIGREQUEST,
    output_type=_INCREMENTALMESHCONFIGRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_AGGREGATEDMESHCONFIGSERVICE)

DESCRIPTOR.services_by_name['AggregatedMeshConfigService'] = _AGGREGATEDMESHCONFIGSERVICE


_RESOURCESOURCE = _descriptor.ServiceDescriptor(
  name='ResourceSource',
  full_name='istio.mcp.v1alpha1.ResourceSource',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  serialized_start=1893,
  serialized_end=2011,
  methods=[
  _descriptor.MethodDescriptor(
    name='EstablishResourceStream',
    full_name='istio.mcp.v1alpha1.ResourceSource.EstablishResourceStream',
    index=0,
    containing_service=None,
    input_type=_REQUESTRESOURCES,
    output_type=_RESOURCES,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_RESOURCESOURCE)

DESCRIPTOR.services_by_name['ResourceSource'] = _RESOURCESOURCE


_RESOURCESINK = _descriptor.ServiceDescriptor(
  name='ResourceSink',
  full_name='istio.mcp.v1alpha1.ResourceSink',
  file=DESCRIPTOR,
  index=2,
  serialized_options=None,
  serialized_start=2013,
  serialized_end=2129,
  methods=[
  _descriptor.MethodDescriptor(
    name='EstablishResourceStream',
    full_name='istio.mcp.v1alpha1.ResourceSink.EstablishResourceStream',
    index=0,
    containing_service=None,
    input_type=_RESOURCES,
    output_type=_REQUESTRESOURCES,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_RESOURCESINK)

DESCRIPTOR.services_by_name['ResourceSink'] = _RESOURCESINK

# @@protoc_insertion_point(module_scope)
