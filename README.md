# dolstats

Package dolstats is basic API to dolstats.com website - "database" of
all processed PERM applications by Department of Labour. It doesn't require
any registration.

Example:
```go
	f := dolstats.Filter{
		Number: caseNumber,
	}
	cases, err := dolstats.GetCases(f)
	if err != nil {
		panic(err)
	}
	if len(cases) == 0 {
		fmt.Println("Sorry, your PERM is still not approved :(")
	}
```
