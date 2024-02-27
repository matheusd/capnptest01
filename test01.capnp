using Go = import "/go.capnp";

@0xf454c62f08bc504b;

$Go.package("capnptest01");
$Go.import("github.com/matheusd/capnptest01");

enum TxType {
	c @0;
	d @1;
}

struct Transaction {
	amount @0 :Int64;
	type @1 :TxType;
	description @2 :Text;
	createdAtMs @3 :Int64;
}
