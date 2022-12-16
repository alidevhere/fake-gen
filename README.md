# US Fake Numbers Generator

`fake-gen` is a single command cli tool. It generates US phone numbers in format `AAA-555-NNNN`. `AAA` is area code, and `NNNN` is any random number between 0001 to 9999.

## Installation:

Linux:
    
    sudo-apt install 

Windows:

    windows install

MacOs:

    brew install

## Use:

|Flag|Short Hand|Use|
|---|---|---|
|column|c|To specify column in output csv file|
|limit|l|To specify rows in output csv file|
|out|o|To specify name of generated file pass only name(no extension) by default it is csv|
|unique|u|To generate unique phone numbers|

Note: 3 types of columns are supported as of now name,age and phone. -c flag can have any of these three values.

## Example-1:

We want to generate a csv file of 10 rows with following configurations 

huge_data.csv

|first_name|last_name|age|mobile_phone|
|---|---|---|---|
|ali|ashraf|22|1234567890|
|...|...|...|...|

We will use command like this:

    fake-gen --limit 10 --out huge_data --column name --column name --column age --column phone

OR by using shorthand

    fake-gen -l 10 -u -o huge_data -c name -c name -c age -c phone



## Example-2:

We want to generate a csv file of 100 rows with following configurations 

huge_data.csv

|first_name|last_name|age|mobile_phone|work_phone|home_phone|
|---|---|---|---|---|---|
|ali|ashraf|22|1234567891|1234567892|1234567893|
|...|...|...|...|

We will use command like this:

    fake-gen --limit 100 --out huge_data --column name --column name --column age --column phone --column phone --column phone

OR by using shorthand

    fake-gen -l 10 -u -o huge_data -c name -c name -c age -c phone -c phone -c phone


