# Repositories

## About
warum, wann, abi, mtune;march, lto, pgo

`x86-64-v3`



## How to add
`$ nano /etc/pkg/dpxlab.conf`
and paste this snippet:
```ucl
dpxlab: {
	url: "https://pkg.dpxlab.de/${ABI}-LATEST/",
	mirror_type: "none",
	signature_type: "pubkey",
	pubkey: "/etc/pkg/hbsdsrv.cert",
	enabled: yes,
	priority: 0,
}
```

- *.conf options described [here](https://man.freebsd.org/cgi/man.cgi?pkg#CONFIGURATION)
- [Public Key of the Repo](/dpxlab.cert)
