# How to reproduce

Simply running `go generate` with Go 1.18beta1 or beta2 will cause the following error:

    2022/02/09 17:47:43 Unexpected package creation during export data loading
    exit status 1

If you comment out the entire policy definitions within `ent/schema/tester.go`. The generator will function. Alternatively, you can leave the privacy block in but comment out the function `GetPermissionsAsList` inside `lib/permissions/bitperm.go` and it will generate. 