# xlsx2sql

![GitHub](https://img.shields.io/github/license/jangabler/xlsx2sql)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jangabler/xlsx2sql)
[![Go Report Card](https://goreportcard.com/badge/github.com/jangabler/xlsx2sql)](https://goreportcard.com/report/github.com/jangabler/xlsx2sql)

Extract data from MS Excel files and generate SQL `INSERT INTO` statements by the use of a declarative mapping.

## Features

* Support for the MS Excel file formats `*.xlsx` and the legacy `*.xls`
* Support for the SQL dialects [MySQL](https://dev.mysql.com/doc/refman/8.0/en/insert.html)/[MariaDB](https://mariadb.com/kb/en/library/insert/) and [PostgreSQL](https://www.postgresql.org/docs/11/sql-insert.html)
* Runs as a command line program on Windows, Linux and macOS

## Use case

* Migrate data from multiple files with the same layout made of MS Excel, e.g. statistics reports, into a SQL database

## Usage

### Mapping

```xml
<?xml version="1.0" encoding="UTF-8"?>
<mapping>
    <xlsx globPattern="spreadsheet*.xlsx">
        <worksheet name="Sheet1">
            <cell coordinate="A1" refAttribute="attr1"/>
            <cell coordinate="B1" refAttribute="attr2"/>
            <cell coordinate="C1" refAttribute="attr3"/>
        </worksheet>
    </xlsx>
    <sql database="db1" table="table1">
        <attribute name="attr1" dataType="string"/>
        <attribute name="attr2" dataType="string"/>
        <attribute name="attr3" dataType="integer"/>
    </sql>
</mapping>
```

### Running

```sh
# Print statements into standard output (stdout)
$ xlsx2sql -m mapping.xml
INSERT INTO db1.table1 (attr1, attr2, attr3) VALUES ('a', 'b', 1);
INSERT INTO db1.table1 (attr1, attr2, attr3) VALUES ('c', 'd', 2);
INSERT INTO db1.table1 (attr1, attr2, attr3) VALUES ('e', 'f', 3);

# Redirect statements into a dump file
$ xlsx2sql -m mapping.xml > mysqldump.sql

# Pipe statements into the MySQL-Client
$ xlsx2sql -m mapping.xml | mysql -u user_name -p your_password db_name
```

## Dependencies

| Package name | License |
| --- | --- |
| [github.com/360EntSecGroup-Skylar/excelize](https://github.com/360EntSecGroup-Skylar/excelize) | [![GitHub license](https://img.shields.io/github/license/360EntSecGroup-Skylar/excelize)](https://github.com/360EntSecGroup-Skylar/excelize/blob/master/LICENSE) |
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | [![GitHub license](https://img.shields.io/github/license/stretchr/testify)](https://github.com/stretchr/testify/blob/master/LICENSE) |
