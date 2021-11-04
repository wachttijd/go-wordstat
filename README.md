# GoWordStat — tool for statistical analysis by words
---
GoWordStat can analyze the amount of use of each word in the texts presenting the result of the analysis in the JSON format.
## Building:
    git clone https://github.com/wachttijd/go-wordstat
    cd go-wordstat
    go build wordstat.go
## Available flags:
    -i Input file for statistical word analysis (default: "input.txt")
    -o Output file for results of analysis (default: "output.json")
    -ignums Ignore numbers (default: false)
    -h Print this help
## Usage:
    ./wordstat -i inputFileWithText.txt -o resultsIn.json
    Preparing text for analysis...
    Text prepared, analysis started...
    Analysis complete!
Now results of analysis `inputFileWithText.txt` are in `resultsIn.json`