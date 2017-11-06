# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hoststore.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import common_pb2 as common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='hoststore.proto',
  package='zsearch',
  syntax='proto3',
  serialized_pb=_b('\n\x0fhoststore.proto\x12\x07zsearch\x1a\x0c\x63ommon.proto\"\x0b\n\tWHOISAtom\"\xf1\x01\n\x0cLocationAtom\x12\x11\n\tcontinent\x18\x01 \x01(\t\x12\x0f\n\x07\x63ountry\x18\x02 \x01(\t\x12\x14\n\x0c\x63ountry_code\x18\x03 \x01(\t\x12\x0c\n\x04\x63ity\x18\x04 \x01(\t\x12\x13\n\x0bpostal_code\x18\x05 \x01(\t\x12\x10\n\x08timezone\x18\x06 \x01(\t\x12\x10\n\x08province\x18\x07 \x01(\t\x12\x10\n\x08latitude\x18\x08 \x01(\x01\x12\x11\n\tlongitude\x18\t \x01(\x01\x12\x1a\n\x12registered_country\x18\n \x01(\t\x12\x1f\n\x17registered_country_code\x18\x0b \x01(\t\"P\n\x0cProtocolAtom\x12$\n\x08metadata\x18\x01 \x03(\x0b\x32\x12.zsearch.Metadatum\x12\x0c\n\x04tags\x18\x02 \x03(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\t\"C\n\x0c\x41nonymousKey\x12\x0c\n\x04port\x18\x01 \x01(\r\x12\x10\n\x08protocol\x18\x02 \x01(\r\x12\x13\n\x0bsubprotocol\x18\x03 \x01(\r\"\xb9\x04\n\x06Record\x12\n\n\x02ip\x18\x01 \x01(\x07\x12\x0c\n\x04port\x18\x02 \x01(\r\x12\x10\n\x08protocol\x18\x03 \x01(\r\x12\x13\n\x0bsubprotocol\x18\x04 \x01(\r\x12\x0e\n\x06\x64omain\x18\x05 \x01(\t\x12\x11\n\ttimestamp\x18\x06 \x01(\x07\x12\x0e\n\x06scanid\x18\x07 \x01(\r\x12\x10\n\x08sha256fp\x18\x08 \x01(\x0c\x12\x1d\n\x15\x66irst_seen_at_scan_id\x18\t \x01(\r\x12\x1c\n\x14last_seen_at_scan_id\x18\n \x01(\r\x12%\n\x04\x61tom\x18\x0b \x01(\x0b\x32\x15.zsearch.ProtocolAtomH\x00\x12\x31\n\x10private_location\x18\x0c \x01(\x0b\x32\x15.zsearch.LocationAtomH\x00\x12\"\n\x07\x61s_atom\x18\r \x01(\x0b\x32\x0f.zsearch.ASAtomH\x00\x12#\n\x05whois\x18\x0e \x01(\x0b\x32\x12.zsearch.WHOISAtomH\x00\x12)\n\x08userdata\x18\x0f \x01(\x0b\x32\x15.zsearch.UserdataAtomH\x00\x12\x30\n\x0fpublic_location\x18\x11 \x01(\x0b\x32\x15.zsearch.LocationAtomH\x00\x12\x14\n\nalexa_rank\x18\x10 \x01(\rH\x00\x12\x18\n\x0equantcast_rank\x18\x13 \x01(\rH\x00\x12\x1d\n\x13\x63isco_umbrella_rank\x18\x14 \x01(\rH\x00\x12\x0f\n\x07version\x18\x12 \x01(\x04\x42\x0c\n\ndata_oneof\"~\n\x05\x44\x65lta\x12&\n\ndelta_type\x18\x01 \x01(\x0e\x32\x12.zsearch.DeltaType\x12\n\n\x02ip\x18\x02 \x01(\x07\x12\x0e\n\x06\x64omain\x18\x03 \x01(\t\x12\x0f\n\x07version\x18\x04 \x01(\x04\x12 \n\x07records\x18\x05 \x03(\x0b\x32\x0f.zsearch.Recordb\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_WHOISATOM = _descriptor.Descriptor(
  name='WHOISAtom',
  full_name='zsearch.WHOISAtom',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=42,
  serialized_end=53,
)


_LOCATIONATOM = _descriptor.Descriptor(
  name='LocationAtom',
  full_name='zsearch.LocationAtom',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='continent', full_name='zsearch.LocationAtom.continent', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='country', full_name='zsearch.LocationAtom.country', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='country_code', full_name='zsearch.LocationAtom.country_code', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='city', full_name='zsearch.LocationAtom.city', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='postal_code', full_name='zsearch.LocationAtom.postal_code', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='timezone', full_name='zsearch.LocationAtom.timezone', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='province', full_name='zsearch.LocationAtom.province', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='latitude', full_name='zsearch.LocationAtom.latitude', index=7,
      number=8, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='longitude', full_name='zsearch.LocationAtom.longitude', index=8,
      number=9, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='registered_country', full_name='zsearch.LocationAtom.registered_country', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='registered_country_code', full_name='zsearch.LocationAtom.registered_country_code', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=56,
  serialized_end=297,
)


_PROTOCOLATOM = _descriptor.Descriptor(
  name='ProtocolAtom',
  full_name='zsearch.ProtocolAtom',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='metadata', full_name='zsearch.ProtocolAtom.metadata', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tags', full_name='zsearch.ProtocolAtom.tags', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='data', full_name='zsearch.ProtocolAtom.data', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=299,
  serialized_end=379,
)


_ANONYMOUSKEY = _descriptor.Descriptor(
  name='AnonymousKey',
  full_name='zsearch.AnonymousKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='port', full_name='zsearch.AnonymousKey.port', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='zsearch.AnonymousKey.protocol', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='subprotocol', full_name='zsearch.AnonymousKey.subprotocol', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=381,
  serialized_end=448,
)


_RECORD = _descriptor.Descriptor(
  name='Record',
  full_name='zsearch.Record',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ip', full_name='zsearch.Record.ip', index=0,
      number=1, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='port', full_name='zsearch.Record.port', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='zsearch.Record.protocol', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='subprotocol', full_name='zsearch.Record.subprotocol', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='domain', full_name='zsearch.Record.domain', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='zsearch.Record.timestamp', index=5,
      number=6, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='scanid', full_name='zsearch.Record.scanid', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sha256fp', full_name='zsearch.Record.sha256fp', index=7,
      number=8, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='first_seen_at_scan_id', full_name='zsearch.Record.first_seen_at_scan_id', index=8,
      number=9, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='last_seen_at_scan_id', full_name='zsearch.Record.last_seen_at_scan_id', index=9,
      number=10, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='atom', full_name='zsearch.Record.atom', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='private_location', full_name='zsearch.Record.private_location', index=11,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='as_atom', full_name='zsearch.Record.as_atom', index=12,
      number=13, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='whois', full_name='zsearch.Record.whois', index=13,
      number=14, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='userdata', full_name='zsearch.Record.userdata', index=14,
      number=15, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='public_location', full_name='zsearch.Record.public_location', index=15,
      number=17, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='alexa_rank', full_name='zsearch.Record.alexa_rank', index=16,
      number=16, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='quantcast_rank', full_name='zsearch.Record.quantcast_rank', index=17,
      number=19, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cisco_umbrella_rank', full_name='zsearch.Record.cisco_umbrella_rank', index=18,
      number=20, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='version', full_name='zsearch.Record.version', index=19,
      number=18, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='data_oneof', full_name='zsearch.Record.data_oneof',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=451,
  serialized_end=1020,
)


_DELTA = _descriptor.Descriptor(
  name='Delta',
  full_name='zsearch.Delta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='delta_type', full_name='zsearch.Delta.delta_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ip', full_name='zsearch.Delta.ip', index=1,
      number=2, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='domain', full_name='zsearch.Delta.domain', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='version', full_name='zsearch.Delta.version', index=3,
      number=4, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='records', full_name='zsearch.Delta.records', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1022,
  serialized_end=1148,
)

_PROTOCOLATOM.fields_by_name['metadata'].message_type = common__pb2._METADATUM
_RECORD.fields_by_name['atom'].message_type = _PROTOCOLATOM
_RECORD.fields_by_name['private_location'].message_type = _LOCATIONATOM
_RECORD.fields_by_name['as_atom'].message_type = common__pb2._ASATOM
_RECORD.fields_by_name['whois'].message_type = _WHOISATOM
_RECORD.fields_by_name['userdata'].message_type = common__pb2._USERDATAATOM
_RECORD.fields_by_name['public_location'].message_type = _LOCATIONATOM
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['atom'])
_RECORD.fields_by_name['atom'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['private_location'])
_RECORD.fields_by_name['private_location'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['as_atom'])
_RECORD.fields_by_name['as_atom'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['whois'])
_RECORD.fields_by_name['whois'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['userdata'])
_RECORD.fields_by_name['userdata'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['public_location'])
_RECORD.fields_by_name['public_location'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['alexa_rank'])
_RECORD.fields_by_name['alexa_rank'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['quantcast_rank'])
_RECORD.fields_by_name['quantcast_rank'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_RECORD.oneofs_by_name['data_oneof'].fields.append(
  _RECORD.fields_by_name['cisco_umbrella_rank'])
_RECORD.fields_by_name['cisco_umbrella_rank'].containing_oneof = _RECORD.oneofs_by_name['data_oneof']
_DELTA.fields_by_name['delta_type'].enum_type = common__pb2._DELTATYPE
_DELTA.fields_by_name['records'].message_type = _RECORD
DESCRIPTOR.message_types_by_name['WHOISAtom'] = _WHOISATOM
DESCRIPTOR.message_types_by_name['LocationAtom'] = _LOCATIONATOM
DESCRIPTOR.message_types_by_name['ProtocolAtom'] = _PROTOCOLATOM
DESCRIPTOR.message_types_by_name['AnonymousKey'] = _ANONYMOUSKEY
DESCRIPTOR.message_types_by_name['Record'] = _RECORD
DESCRIPTOR.message_types_by_name['Delta'] = _DELTA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

WHOISAtom = _reflection.GeneratedProtocolMessageType('WHOISAtom', (_message.Message,), dict(
  DESCRIPTOR = _WHOISATOM,
  __module__ = 'hoststore_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.WHOISAtom)
  ))
_sym_db.RegisterMessage(WHOISAtom)

LocationAtom = _reflection.GeneratedProtocolMessageType('LocationAtom', (_message.Message,), dict(
  DESCRIPTOR = _LOCATIONATOM,
  __module__ = 'hoststore_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.LocationAtom)
  ))
_sym_db.RegisterMessage(LocationAtom)

ProtocolAtom = _reflection.GeneratedProtocolMessageType('ProtocolAtom', (_message.Message,), dict(
  DESCRIPTOR = _PROTOCOLATOM,
  __module__ = 'hoststore_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.ProtocolAtom)
  ))
_sym_db.RegisterMessage(ProtocolAtom)

AnonymousKey = _reflection.GeneratedProtocolMessageType('AnonymousKey', (_message.Message,), dict(
  DESCRIPTOR = _ANONYMOUSKEY,
  __module__ = 'hoststore_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.AnonymousKey)
  ))
_sym_db.RegisterMessage(AnonymousKey)

Record = _reflection.GeneratedProtocolMessageType('Record', (_message.Message,), dict(
  DESCRIPTOR = _RECORD,
  __module__ = 'hoststore_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.Record)
  ))
_sym_db.RegisterMessage(Record)

Delta = _reflection.GeneratedProtocolMessageType('Delta', (_message.Message,), dict(
  DESCRIPTOR = _DELTA,
  __module__ = 'hoststore_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.Delta)
  ))
_sym_db.RegisterMessage(Delta)


# @@protoc_insertion_point(module_scope)
