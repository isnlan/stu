// This file is generated. Do not edit
// @generated

// https://github.com/Manishearth/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![cfg_attr(rustfmt, rustfmt_skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unsafe_code)]
#![allow(unused_imports)]
#![allow(unused_results)]

const METHOD_KV_DB_SET: ::grpcio::Method<super::kvdb::SetKV, super::kvdb::Empty> = ::grpcio::Method {
    ty: ::grpcio::MethodType::Unary,
    name: "/KvDB/Set",
    req_mar: ::grpcio::Marshaller { ser: ::grpcio::pb_ser, de: ::grpcio::pb_de },
    resp_mar: ::grpcio::Marshaller { ser: ::grpcio::pb_ser, de: ::grpcio::pb_de },
};

#[derive(Clone)]
pub struct KvDbClient {
    client: ::grpcio::Client,
}

impl KvDbClient {
    pub fn new(channel: ::grpcio::Channel) -> Self {
        KvDbClient {
            client: ::grpcio::Client::new(channel),
        }
    }

    pub fn set_opt(&self, req: &super::kvdb::SetKV, opt: ::grpcio::CallOption) -> ::grpcio::Result<super::kvdb::Empty> {
        self.client.unary_call(&METHOD_KV_DB_SET, req, opt)
    }

    pub fn set(&self, req: &super::kvdb::SetKV) -> ::grpcio::Result<super::kvdb::Empty> {
        self.set_opt(req, ::grpcio::CallOption::default())
    }

    pub fn set_async_opt(&self, req: &super::kvdb::SetKV, opt: ::grpcio::CallOption) -> ::grpcio::Result<::grpcio::ClientUnaryReceiver<super::kvdb::Empty>> {
        self.client.unary_call_async(&METHOD_KV_DB_SET, req, opt)
    }

    pub fn set_async(&self, req: &super::kvdb::SetKV) -> ::grpcio::Result<::grpcio::ClientUnaryReceiver<super::kvdb::Empty>> {
        self.set_async_opt(req, ::grpcio::CallOption::default())
    }
    pub fn spawn<F>(&self, f: F) where F: ::futures::Future<Item = (), Error = ()> + Send + 'static {
        self.client.spawn(f)
    }
}

pub trait KvDb {
    fn set(&mut self, ctx: ::grpcio::RpcContext, req: super::kvdb::SetKV, sink: ::grpcio::UnarySink<super::kvdb::Empty>);
}

pub fn create_kv_db<S: KvDb + Send + Clone + 'static>(s: S) -> ::grpcio::Service {
    let mut builder = ::grpcio::ServiceBuilder::new();
    let mut instance = s.clone();
    builder = builder.add_unary_handler(&METHOD_KV_DB_SET, move |ctx, req, resp| {
        instance.set(ctx, req, resp)
    });
    builder.build()
}
