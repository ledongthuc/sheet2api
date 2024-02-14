# Sheet2API

## License

The source code is under GPL 3.0 and free for non-commercial use.

If you need it for commercial usage, feel free to contact me and it's only $1 per month.

## Quickstart

1. Copy `config.yaml.example` to `config.yaml`

2. Run script

```bash
go run main.go
```

The server will start default with binding IP 0.0.0.0 and port 14119 with following logs:

```bash
$ go run main.go

INF Config file 'config.yaml'
Start server: 0.0.0.0:14119
```

3. Access URL http://localhost:14119/test/users and get the response

```json
[
  {
    "Age": "32",
    "Country": "United States",
    "Date": "15/10/2017",
    "First Name": "Dulce",
    "Gender": "Female",
    "Id": "1562",
    "Last Name": "Abril"
  },
  {
    "Age": "25",
    "Country": "Great Britain",
    "Date": "16/08/2016",
    "First Name": "Mara",
    "Gender": "Female",
    "Id": "1582",
    "Last Name": "Hashimoto"
  },
  {
    "Age": "36",
    "Country": "France",
    "Date": "21/05/2015",
    "First Name": "Philip",
    "Gender": "Male",
    "Id": "2587",
    "Last Name": "Gent"
  }
]
```

## TODO

 - Support sheet files content caching
 - Support sheet column name/API response key replaced name
 - Rename package to SheetFetch
 - Document config
