// Autogenerated by Frugal Compiler (0.0.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

library event.src.f_foo_scope;

import 'dart:async';

import 'dart:typed_data' show Uint8List;
import 'package:thrift/thrift.dart' as thrift;
import 'package:frugal/frugal.dart' as frugal;

import 'event.dart' as t_event;
import 'foo.dart' as t_foo;


abstract class FFoo {
  // Ping the server.
  Future ping(frugal.Context ctx);
  // Blah the server.
  Future<int> blah(frugal.Context ctx, int num, String str, t_event.Event event);
}

class FFooClient implements FFoo {

  FFooClient(thrift.TProtocol iprot, [thrift.TProtocol oprot = null]) {
    _iprot = iprot;
    _oprot = (oprot == null) ? iprot : oprot;
  }

  frugal.FProtocol _iprot;

  frugal.FProtocol get iprot => _iprot;

  frugal.FProtocol _oprot;

  frugal.FProtocol get oprot => _oprot;

  int _seqid = 0;

  int get seqid => _seqid;

  int nextSeqid() => ++_seqid;

  // Ping the server.
  Future ping(frugal.Context ctx) async {
    oprot.writeRequestHeader(ctx);
    oprot.writeMessageBegin(new thrift.TMessage("ping", thrift.TMessageType.CALL, nextSeqid()));
    t_foo.ping_args args = new t_foo.ping_args();
    args.write(oprot);
    oprot.writeMessageEnd();

    await oprot.transport.flush();

    iprot.readResponseHeader(ctx);
    thrift.TMessage msg = iprot.readMessageBegin();
    if (msg.type == thrift.TMessageType.EXCEPTION) {
      thrift.TApplicationError error = thrift.TApplicationError.read(iprot);
      iprot.readMessageEnd();
      throw error;
    }

    t_foo.ping_result result = new t_foo.ping_result();
    result.read(iprot);
    iprot.readMessageEnd();
    return;
  }

  // Blah the server.
  Future<int> blah(frugal.Context ctx, int num, String str, t_event.Event event) async {
    oprot.writeRequestHeader(ctx);
    oprot.writeMessageBegin(new thrift.TMessage("blah", thrift.TMessageType.CALL, nextSeqid()));
    t_foo.blah_args args = new t_foo.blah_args();
    args.num = num;
    args.str = str;
    args.event = event;
    args.write(oprot);
    oprot.writeMessageEnd();

    await oprot.transport.flush();

    iprot.readResponseHeader(ctx);
    thrift.TMessage msg = iprot.readMessageBegin();
    if (msg.type == thrift.TMessageType.EXCEPTION) {
      thrift.TApplicationError error = thrift.TApplicationError.read(iprot);
      iprot.readMessageEnd();
      throw error;
    }

    t_foo.blah_result result = new t_foo.blah_result();
    result.read(iprot);
    iprot.readMessageEnd();
    if (result.isSetSuccess()) {
      return result.success;
    }

    throw new thrift.TApplicationError(thrift.TApplicationErrorType.MISSING_RESULT, "blah failed: unknown result");
  }

}
