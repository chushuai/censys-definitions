# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protocols.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='protocols.proto',
  package='zsearch',
  syntax='proto3',
  serialized_pb=_b('\n\x0fprotocols.proto\x12\x07zsearch*\xfb\x05\n\x08Protocol\x12\x12\n\x0ePROTO_RESERVED\x10\x00\x12\x10\n\x0cPROTO_SYSTEM\x10\x01\x12\x0e\n\nPROTO_HTTP\x10\x02\x12\x0f\n\x0bPROTO_HTTPS\x10\x03\x12\x0e\n\nPROTO_IMAP\x10\x04\x12\x0f\n\x0bPROTO_IMAPS\x10\x05\x12\x0e\n\nPROTO_SMTP\x10\x06\x12\x0f\n\x0bPROTO_SMTPS\x10\x07\x12\x0e\n\nPROTO_POP3\x10\x08\x12\x0f\n\x0bPROTO_POP3S\x10\t\x12\x10\n\x0cPROTO_MODBUS\x10\n\x12\r\n\tPROTO_FTP\x10\x0b\x12\r\n\tPROTO_SSH\x10\x0c\x12\r\n\tPROTO_DNS\x10\r\x12\r\n\tPROTO_NTP\x10\x0e\x12\x10\n\x0cPROTO_TELNET\x10\x0f\x12\x0e\n\nPROTO_UPNP\x10\x10\x12\x0e\n\nPROTO_CWMP\x10\x11\x12\x10\n\x0cPROTO_HTTP_2\x10\x12\x12\x10\n\x0cPROTO_BACNET\x10\x13\x12\x0e\n\nPROTO_DNP3\x10\x14\x12\r\n\tPROTO_FOX\x10\x15\x12\x0c\n\x08PROTO_S7\x10\x16\x12\x10\n\x0cPROTO_GLOBAL\x10\x17\x12\x10\n\x0cPROTO_LOOKUP\x10\x18\x12\x12\n\x0ePROTO_HTTP_WWW\x10\x1a\x12\x13\n\x0fPROTO_HTTPS_WWW\x10\x1b\x12\r\n\tPROTO_SMB\x10\x1c\x12\x10\n\x0bPROTO_MYSQL\x10\xad\x02\x12\x10\n\x0bPROTO_REDIS\x10\xae\x02\x12\x0e\n\tPROTO_RDP\x10\xaf\x02\x12\x13\n\x0ePROTO_MEMCACHE\x10\xb0\x02\x12\x12\n\rPROTO_MONGODB\x10\xb1\x02\x12\x0f\n\nPROTO_AMQP\x10\xb2\x02\x12\x0f\n\nPROTO_RIAK\x10\xb3\x02\x12\x0e\n\tPROTO_IKE\x10\xb4\x02\x12\x12\n\rPROTO_CODESYS\x10\xb5\x02\x12\x11\n\x0cPROTO_PCWORX\x10\xb6\x02\x12\x10\n\x0bPROTO_OMRON\x10\xb7\x02\x12\x15\n\x10PROTO_ETHERNETIP\x10\xb8\x02\x12\x10\n\x0bPROTO_PGSQL\x10\xb9\x02\x12\x12\n\rPROTO_MELSECQ\x10\xba\x02\x12\x13\n\x0ePROTO_PROCONOS\x10\xbb\x02*\xe7\n\n\x0bSubprotocol\x12\x15\n\x11SUBPROTO_RESERVED\x10\x00\x12\x14\n\x10SUBPROTO_DELETED\x10\x01\x12\x14\n\x10SUBPROTO_GENERIC\x10\x02\x12\x13\n\x0fSUBPROTO_BANNER\x10\x03\x12\x10\n\x0cSUBPROTO_TLS\x10\x04\x12\x14\n\x10SUBPROTO_TLS_1_0\x10\x05\x12\x14\n\x10SUBPROTO_TLS_1_1\x10\x06\x12\x14\n\x10SUBPROTO_TLS_1_2\x10\x07\x12\x14\n\x10SUBPROTO_TLS_1_3\x10\x08\x12\x17\n\x13SUBPROTO_HEARTBLEED\x10\t\x12\x14\n\x10SUBPROTO_CIPHERS\x10\n\x12\x12\n\x0eSUBPROTO_SSL_2\x10\x0b\x12\x12\n\x0eSUBPROTO_SSL_3\x10\x0c\x12\x10\n\x0cSUBPROTO_GET\x10\r\x12\x15\n\x11SUBPROTO_STARTTLS\x10\x0e\x12\x13\n\x0fSUBPROTO_EXPORT\x10\x0f\x12\x17\n\x13SUBPROTO_RSA_EXPORT\x10\x10\x12\x17\n\x13SUBPROTO_DHE_EXPORT\x10\x11\x12\x10\n\x0cSUBPROTO_DHE\x10\x12\x12\x12\n\x0eSUBPROTO_ECDHE\x10\x13\x12\x10\n\x0cSUBPROTO_SNI\x10\x14\x12\x13\n\x0fSUBPROTO_NO_SNI\x10\x15\x12\x11\n\rSUBPROTO_QUIC\x10\x16\x12\x11\n\rSUBPROTO_SPDY\x10\x17\x12\x10\n\x0cSUBPROTO_RSA\x10\x18\x12\x10\n\x0cSUBPROTO_DSA\x10\x19\x12\x12\n\x0eSUBPROTO_ECDSA\x10\x1a\x12\x16\n\x12SUBPROTO_DEVICE_ID\x10\x1b\x12\x1a\n\x16SUBPROTO_OPEN_RESOLVER\x10\x1c\x12\x17\n\x13SUBPROTO_OPEN_PROXY\x10\x1d\x12\x17\n\x13SUBPROTO_OPEN_RELAY\x10\x1e\x12\x11\n\rSUBPROTO_TIME\x10\x1f\x12\x19\n\x15SUBPROTO_HACKING_TEAM\x10 \x12\x1c\n\x18SUBPROTO_EXTENDED_RANDOM\x10!\x12\x16\n\x12SUBPROTO_DISCOVERY\x10\"\x12\x13\n\x0fSUBPROTO_GTLD_A\x10#\x12\x13\n\x0fSUBPROTO_LOOKUP\x10$\x12\x13\n\x0fSUBPROTO_STATUS\x10%\x12\x10\n\x0cSUBPROTO_SZL\x10&\x12\x0f\n\x0bSUBPROTO_V2\x10\'\x12!\n\x1cSUBPROTO_SYS_PUBLIC_LOCATION\x10\xc0\x01\x12\x14\n\x0fSUBPROTO_SYS_AS\x10\xc1\x01\x12\x16\n\x11SUBPROTO_SYS_TAGS\x10\xc2\x01\x12\x1a\n\x15SUBPROTO_SYS_METADATA\x10\xc3\x01\x12\x17\n\x12SUBPROTO_SYS_WHOIS\x10\xc4\x01\x12\x1a\n\x15SUBPROTO_SYS_USERDATA\x10\xc5\x01\x12\x1b\n\x16SUBPROTO_SYS_BLACKLIST\x10\xc6\x01\x12\x1c\n\x17SUBPROTO_SYS_ALEXA_RANK\x10\xc7\x01\x12%\n SUBPROTO_SYS_RESTRICTED_LOCATION\x10\xc8\x01\x12\x19\n\x14SUBPROTO_SYS_VERSION\x10\xc9\x01\x12 \n\x1bSUBPROTO_SYS_QUANTCAST_RANK\x10\xca\x01\x12%\n SUBPROTO_SYS_CISCO_UMBRELLA_RANK\x10\xcb\x01\x12\x1d\n\x18SUBPROTO_SYS_REVERSE_DNS\x10\xcc\x01\x12\x11\n\x0cSUBPROTO_SPF\x10\xdc\x01\x12\x13\n\x0eSUBPROTO_DMARC\x10\xdd\x01\x12\x12\n\rSUBPROTO_DKIM\x10\xde\x01\x12\x0f\n\nSUBPROTO_A\x10\xdf\x01\x12\x10\n\x0bSUBPROTO_MX\x10\xe0\x01\x12\x12\n\rSUBPROTO_AXFR\x10\xe1\x01\x62\x06proto3')
)

_PROTOCOL = _descriptor.EnumDescriptor(
  name='Protocol',
  full_name='zsearch.Protocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PROTO_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_SYSTEM', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_HTTP', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_HTTPS', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_IMAP', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_IMAPS', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_SMTP', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_SMTPS', index=7, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_POP3', index=8, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_POP3S', index=9, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_MODBUS', index=10, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_FTP', index=11, number=11,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_SSH', index=12, number=12,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_DNS', index=13, number=13,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_NTP', index=14, number=14,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_TELNET', index=15, number=15,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_UPNP', index=16, number=16,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_CWMP', index=17, number=17,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_HTTP_2', index=18, number=18,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_BACNET', index=19, number=19,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_DNP3', index=20, number=20,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_FOX', index=21, number=21,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_S7', index=22, number=22,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_GLOBAL', index=23, number=23,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_LOOKUP', index=24, number=24,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_HTTP_WWW', index=25, number=26,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_HTTPS_WWW', index=26, number=27,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_SMB', index=27, number=28,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_MYSQL', index=28, number=301,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_REDIS', index=29, number=302,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_RDP', index=30, number=303,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_MEMCACHE', index=31, number=304,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_MONGODB', index=32, number=305,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_AMQP', index=33, number=306,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_RIAK', index=34, number=307,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_IKE', index=35, number=308,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_CODESYS', index=36, number=309,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_PCWORX', index=37, number=310,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_OMRON', index=38, number=311,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_ETHERNETIP', index=39, number=312,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_PGSQL', index=40, number=313,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_MELSECQ', index=41, number=314,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROTO_PROCONOS', index=42, number=315,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=29,
  serialized_end=792,
)
_sym_db.RegisterEnumDescriptor(_PROTOCOL)

Protocol = enum_type_wrapper.EnumTypeWrapper(_PROTOCOL)
_SUBPROTOCOL = _descriptor.EnumDescriptor(
  name='Subprotocol',
  full_name='zsearch.Subprotocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DELETED', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_GENERIC', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_BANNER', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_TLS', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_TLS_1_0', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_TLS_1_1', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_TLS_1_2', index=7, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_TLS_1_3', index=8, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_HEARTBLEED', index=9, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_CIPHERS', index=10, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SSL_2', index=11, number=11,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SSL_3', index=12, number=12,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_GET', index=13, number=13,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_STARTTLS', index=14, number=14,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_EXPORT', index=15, number=15,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_RSA_EXPORT', index=16, number=16,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DHE_EXPORT', index=17, number=17,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DHE', index=18, number=18,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_ECDHE', index=19, number=19,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SNI', index=20, number=20,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_NO_SNI', index=21, number=21,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_QUIC', index=22, number=22,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SPDY', index=23, number=23,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_RSA', index=24, number=24,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DSA', index=25, number=25,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_ECDSA', index=26, number=26,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DEVICE_ID', index=27, number=27,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_OPEN_RESOLVER', index=28, number=28,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_OPEN_PROXY', index=29, number=29,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_OPEN_RELAY', index=30, number=30,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_TIME', index=31, number=31,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_HACKING_TEAM', index=32, number=32,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_EXTENDED_RANDOM', index=33, number=33,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DISCOVERY', index=34, number=34,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_GTLD_A', index=35, number=35,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_LOOKUP', index=36, number=36,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_STATUS', index=37, number=37,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SZL', index=38, number=38,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_V2', index=39, number=39,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_PUBLIC_LOCATION', index=40, number=192,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_AS', index=41, number=193,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_TAGS', index=42, number=194,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_METADATA', index=43, number=195,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_WHOIS', index=44, number=196,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_USERDATA', index=45, number=197,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_BLACKLIST', index=46, number=198,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_ALEXA_RANK', index=47, number=199,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_RESTRICTED_LOCATION', index=48, number=200,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_VERSION', index=49, number=201,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_QUANTCAST_RANK', index=50, number=202,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_CISCO_UMBRELLA_RANK', index=51, number=203,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SYS_REVERSE_DNS', index=52, number=204,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_SPF', index=53, number=220,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DMARC', index=54, number=221,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_DKIM', index=55, number=222,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_A', index=56, number=223,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_MX', index=57, number=224,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBPROTO_AXFR', index=58, number=225,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=795,
  serialized_end=2178,
)
_sym_db.RegisterEnumDescriptor(_SUBPROTOCOL)

Subprotocol = enum_type_wrapper.EnumTypeWrapper(_SUBPROTOCOL)
PROTO_RESERVED = 0
PROTO_SYSTEM = 1
PROTO_HTTP = 2
PROTO_HTTPS = 3
PROTO_IMAP = 4
PROTO_IMAPS = 5
PROTO_SMTP = 6
PROTO_SMTPS = 7
PROTO_POP3 = 8
PROTO_POP3S = 9
PROTO_MODBUS = 10
PROTO_FTP = 11
PROTO_SSH = 12
PROTO_DNS = 13
PROTO_NTP = 14
PROTO_TELNET = 15
PROTO_UPNP = 16
PROTO_CWMP = 17
PROTO_HTTP_2 = 18
PROTO_BACNET = 19
PROTO_DNP3 = 20
PROTO_FOX = 21
PROTO_S7 = 22
PROTO_GLOBAL = 23
PROTO_LOOKUP = 24
PROTO_HTTP_WWW = 26
PROTO_HTTPS_WWW = 27
PROTO_SMB = 28
PROTO_MYSQL = 301
PROTO_REDIS = 302
PROTO_RDP = 303
PROTO_MEMCACHE = 304
PROTO_MONGODB = 305
PROTO_AMQP = 306
PROTO_RIAK = 307
PROTO_IKE = 308
PROTO_CODESYS = 309
PROTO_PCWORX = 310
PROTO_OMRON = 311
PROTO_ETHERNETIP = 312
PROTO_PGSQL = 313
PROTO_MELSECQ = 314
PROTO_PROCONOS = 315
SUBPROTO_RESERVED = 0
SUBPROTO_DELETED = 1
SUBPROTO_GENERIC = 2
SUBPROTO_BANNER = 3
SUBPROTO_TLS = 4
SUBPROTO_TLS_1_0 = 5
SUBPROTO_TLS_1_1 = 6
SUBPROTO_TLS_1_2 = 7
SUBPROTO_TLS_1_3 = 8
SUBPROTO_HEARTBLEED = 9
SUBPROTO_CIPHERS = 10
SUBPROTO_SSL_2 = 11
SUBPROTO_SSL_3 = 12
SUBPROTO_GET = 13
SUBPROTO_STARTTLS = 14
SUBPROTO_EXPORT = 15
SUBPROTO_RSA_EXPORT = 16
SUBPROTO_DHE_EXPORT = 17
SUBPROTO_DHE = 18
SUBPROTO_ECDHE = 19
SUBPROTO_SNI = 20
SUBPROTO_NO_SNI = 21
SUBPROTO_QUIC = 22
SUBPROTO_SPDY = 23
SUBPROTO_RSA = 24
SUBPROTO_DSA = 25
SUBPROTO_ECDSA = 26
SUBPROTO_DEVICE_ID = 27
SUBPROTO_OPEN_RESOLVER = 28
SUBPROTO_OPEN_PROXY = 29
SUBPROTO_OPEN_RELAY = 30
SUBPROTO_TIME = 31
SUBPROTO_HACKING_TEAM = 32
SUBPROTO_EXTENDED_RANDOM = 33
SUBPROTO_DISCOVERY = 34
SUBPROTO_GTLD_A = 35
SUBPROTO_LOOKUP = 36
SUBPROTO_STATUS = 37
SUBPROTO_SZL = 38
SUBPROTO_V2 = 39
SUBPROTO_SYS_PUBLIC_LOCATION = 192
SUBPROTO_SYS_AS = 193
SUBPROTO_SYS_TAGS = 194
SUBPROTO_SYS_METADATA = 195
SUBPROTO_SYS_WHOIS = 196
SUBPROTO_SYS_USERDATA = 197
SUBPROTO_SYS_BLACKLIST = 198
SUBPROTO_SYS_ALEXA_RANK = 199
SUBPROTO_SYS_RESTRICTED_LOCATION = 200
SUBPROTO_SYS_VERSION = 201
SUBPROTO_SYS_QUANTCAST_RANK = 202
SUBPROTO_SYS_CISCO_UMBRELLA_RANK = 203
SUBPROTO_SYS_REVERSE_DNS = 204
SUBPROTO_SPF = 220
SUBPROTO_DMARC = 221
SUBPROTO_DKIM = 222
SUBPROTO_A = 223
SUBPROTO_MX = 224
SUBPROTO_AXFR = 225


DESCRIPTOR.enum_types_by_name['Protocol'] = _PROTOCOL
DESCRIPTOR.enum_types_by_name['Subprotocol'] = _SUBPROTOCOL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)


# @@protoc_insertion_point(module_scope)
