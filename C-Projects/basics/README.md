# Compilation using GCC (GNU Compiler Collection)

## Naive Way
Simplist way is to simply call:
```
gcc example.c
```
This will output to a.out.
Which can be run like an executable or script from the terminal with
```
./a.out
```

## Single File
To specify the output name simply use the -o "output" option.
```
gcc -o example.exe example.c
```
Some extra useful options to pass include:
* -Wall which will show all warnings.
* -g which will output extra debugging options for gdb debuggers (vscode c/c++ extension has one such debugger).

## Multiple Files
Multiple files will need to be linked together in order to output correctly. Gcc is smart enough to do this itself however it can be done manually in multiple steps if the user so sees fit.
This will use the -c argument to only compile into object files and then link the output files together into an executable.
```
// Compile Only
gcc -c -Wall -g example.c example2.c ...
// Link Files
gcc -g -o example.exe example.o example2.o ...
```
To do them together simply call gcc in the same way you would compile one file but with multiple files listed.
However in order to do this one must create header files for the imported modules and import them into the file containing int main.
```
gcc -o example.exe example.c example2.c
```
You may also include the -Wall and -g functions if that wasn't clear