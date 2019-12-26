# scoop-uninstall-with-deps

Simple program that uninstalls a [scoop](https://scoop.sh) package, along with its dependencies.

## Usage
``go build`` it (``scoop install go``), and optionally add this function to your PowerShell profile:

```powershell
function scoop {
  if ( $args[0] -eq "uninstall" ) {
    $first, $rest= $args
    Invoke-Expression "path\to\scoop-uninstall-with-deps.exe $rest"
  } else {
    Invoke-Expression "C:\Users\<your user>\scoop\apps\scoop\current\bin\scoop.ps1 $args"
  }
}
```