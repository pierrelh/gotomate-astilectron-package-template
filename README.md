# Gotomate-Astilectron Package Template

This is a template of a Gotomate-Astilectron's package that you can use to create your own Gotomate-Astilectron's package.

You can choose to create your own functions to do whatever you want with a Gotomate-Astilectron's process.

## Installation

Download or Clone the repo:
```
git pull https://github.com/pierrelh/gotomate-astilectron-package-template.git
```

## Setup

Fill the all the files in the Package folder with the right requirements. Fill free to use as an exemple the originals Gotomate-Astilectron's packages to help you building this.

Change the icon.png with an icon corresponding to your package for a better use by users. The size of the icon must be 32x32px. 

Then rename the "Package" folder to what your package should be named (don't forget to put the folder name on uppercase)

## Installation of the package

### Installation by Gotomate-Astilectron

In Gotomate-Astilectron's action menu, select "Import package" then select your .zip archive & restart Gotomate-Astilectron. If the package was correctly
builded then it should work !

### Manually installation

#### Step 1

Copy & Paste your folder into fiber/packages directory

#### Step 2

Go to fiber/fiber.go & add this:
```
case "YourFolderName":
	nextID = yourPackageNamepack.Processing(funcName, instructionData, finished)
```
after the line 229 (at the time I'm writing this), were all the packages are called for the processing.
Be sure that your package is imported at the top of the file and that the import is called yourPackageName + pack.

#### Step 3

Go to fiber/packages/packages-dialog.go & add this:
```
case "YourFolderName":
	databinder, template = yourPackageNamepack.Build(funcName)
```
after the line 42 (at the time I'm writing this), were all the packages are called for the building.
Be sure that your package is imported at the top of the file and that the import is called yourPackageName + pack.


Now the package should be implemented !