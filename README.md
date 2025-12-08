# Yaa
Yaa — YAML Search for Humans

Yaa is a CLI for fast, full‑text search across YAML‑based projects. It’s great for detection content (SigmaHQ, Splunk, Nuclei, Sentinel) and any repo that organizes knowledge as YAML. It builds a local Bleve index and lets you query with a simple, expressive language.

Index location: by default Yaa creates a `yaml_index` folder in your current working directory.

## Query Language 
The power of yaa search comes from its query language, which depends on full-text search. Search can be as simple as writing a few keywords to search for any matches or to search inside a specific yaml property, or both. 

- +keywords means the word must appear (**AND** operation). When using multiple keywords, all specified keywords must appear in the search results.
- -keywords means the keyword must not appear in the search result (**Not** operation). When using multiple keywords, none of the specified keywords should appear in the search results.

Examples:

- Full text: `7zip`
- AND: `+powershell +obfuscation`
- NOT: `wmi -falsepositive:wmi`
- Field search: `title:"powershell obfuscation"`
- Nested field: `.metadata.author:"alice"`


### Full Text Search

Search all rules that contain the keyword '7zip'. 

![image](https://github.com/alwashali/yaa/assets/22593441/85a9905c-6bb2-44e3-9e33-9b05f107322d)

 
Yaml property search with AND operator, searching for any rule that has the word 'powershell **AND** obfuscation' in the title property.

![image](https://github.com/alwashali/yaa/assets/22593441/cb1ba680-b539-459d-92f7-b0f5e4317824)


### Exclude Result

Search for **WMI** persistence related rules and exclude any rule having **WMI** in the falsepositive property. 

![image](https://github.com/alwashali/yaa/assets/22593441/8007a61b-7b91-483f-b330-b5ea45c336a8)


### Nested property search

Yaa can index nested properties and make them searchable by specifying the nested property name prefixed with a dot.

![image](https://github.com/alwashali/yaa/assets/22593441/b5ea4e28-b481-4277-b308-7d0b536b1d69)



### Export Matches

Files matching the search criteria can be exported to a different directory

 ![image](https://github.com/alwashali/yaa/assets/22593441/ca5f6433-0b24-4ad7-b495-26bd67ff8354)


Note: export copies files by basename only and does not preserve directory structure. Use `--force` to overwrite existing files.


### Indexing yaml project 

yaa is built for searching inside detection rule projects such as SigmaHQ, however it can be used with any similar project. To index a yaml project, use the command **index**. 


```
% git clone https://github.com/SigmaHQ/sigma.git
% ./yaa index sigma/rules/
```

![image](https://github.com/alwashali/yaa/assets/22593441/886d03f6-2120-4d22-a5e2-4530a68bf018)


## Installation

Build from source with Go:

```bash
% git clone https://github.com/alwashali/yaa.git
% go build -o yaa
% ./yaa
```

## Usage

Commands:

- `index, i`: build/update the local index
- `search, s`: query the index and optionally export matches

### Index

Synopsis: `yaa index [options] <folder>`

Options:
- `--debug, -d`: enable verbose debug logging

Example:

```bash
./yaa index -d ./sigma/rules
```

### Search

Synopsis: `yaa search [options] <query...>`

Options:
- `--limit, -l`: number of results to display (default: 10)
- `--export, -e`: path to save matched YAML files
- `--force, -f`: overwrite existing files when exporting
- `--debug, -d`: enable verbose debug logging

Examples:

```bash
# Simple keyword
./yaa search "7zip"

# AND and field search
./yaa search -l 5 "+powershell +obfuscation title:obfuscation"

# Exclude matches and export
./yaa search -e /tmp/export -f "wmi -falsepositive:wmi"
```

## Troubleshooting

- "Index was not found": run `yaa index <folder>` first and ensure you are in the same working directory where `yaml_index` exists.
- Empty results: simplify the query, check field names, and try removing `-keyword` filters.
- Export errors: verify destination path; use `--force` to overwrite conflicts.

## License

MIT (see `LICENSE` if present).





