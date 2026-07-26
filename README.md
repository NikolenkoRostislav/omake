# Omake

Omake is a small wrapper around Make that lets you run your Makefile targets from anywhere on your system.

## Why I made Omake

Makefiles are great for automating commands, but they are usually tied to a specific project directory. I wanted a way to use them globally so I made Omake!

Omake adds a global configuration layer that lets you use Make targets as personal commands.

Instead of writing:

```bash
make -f ~/path/to/Makefile -C ~/project target
```

you can simply run:

```bash
omake target
```

Omake handles the Makefile path, execution directory, variables, and provides extra features like listing and describing targets.

## How to use Omake

Use ```omake setup config``` to initialize the omake config inside your home folder. 

You can choose the config location using ```omake setup config <config directory>```. 

A folder will be created with a ```config.yaml``` and a ```Makefile```. 

Please use ```omake setup config <config directory>``` if you want to change the config's location.

An example [config.yaml](./example/config.yaml) and [Makefile](./example/Makefile) are provided. Let's go over the configuration options:

```yaml
targets:
  glog:
    ...

  gsave:
    ...

  gsave-new:
    ...

  gnew-branch:
    ...

  extract-audio:
    ...

  drestart:
    ...

```

Every Makefile target that you want to use with omake must be listed under targets. A target can have the following fields:

### 1. Descriptions

```yaml
  glog:
    description: "logs all git commits"
```

The description will be displayed when running ```omake describe <target>``` 

This field is optional.

### 2. Execution Directory

```yaml
  gsave:
    execution_dir: "."
```
The directory where the Make target will be executed.

If not specified, it defaults to the current directory (```.```).

You can use:
- ```~``` for your home directory
- Any valid path, for example ```D:/Work/omake```

### 3. Variables

```yaml
  gnew-branch:
    variables:
      - name: new
        description: "Name of the new branch"
        env_var: "NEW_BRANCH"
      - name: source
        env_var: "SOURCE_BRANCH"
        default: "main"
```
Variables define values that are passed to the Makefile.

Each variable requires:
- name - the variable name used by omake
- env_var - the environment variable passed to Make

Optional fields:
- description - shown in omake describe <target>
- default - value used when no value is provided

Variable rules
- Variable order in the config determines positional argument order.
- Required variables must appear before variables with defaults.
- Positional arguments cannot be used after keyword arguments.

## How to use variables

Omake allows for the use of positional arguments:

```yaml
omake <target> <argument 1> <argument 2>
```

You can also use variable names

```yaml
omake <target> something="some value" something-else="a different value"
```

Positional and keyword arguments can be combined. Keyword arguments must always come after positional arguments:

```yaml
omake <target> "some value" some-variable="42"
```

## Installation

Clone the repository and run:
```bash
make
```
from the project directory to build the binary.

The compiled executable will be available at `./bin/omake.exe`.

## Notes

- Make sure to add omake to your PATH after installation.
- Currently, omake has only been tested on Windows. Linux support is planned for the future.

If you have any questions or feedback about the project, feel free to open an issue or reach out.