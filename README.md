# eth-contract-address
Contract addresses in ethereum are deterministic. This small command line utility calculates a contract address for a given address and nonce

Usage:

    go run eth-contract-address.go -address=0x6ac7ea33f8831ea9dcc53393aaa88b25a785dbf0 -nonce=0
    # 0xcd234A471b72ba2F1Ccf0A70FCABA648a5eeCD8d
    go run eth-contract-address.go -address=0x6ac7ea33f8831ea9dcc53393aaa88b25a785dbf0 -nonce=1
    # 0x343c43A37D37dfF08AE8C4A11544c718AbB4fCF8
