# Python Virtual Environment Experiments

Setting up a Python virtual environment.

## Create environment

Use `python3 -m venv .venv` to create a virtual environment in the `.venv` directory.

## Configure PyRight

Create a `pyrightconfig.json` file with the following content:

```JSON
{
  "venvPath": ".",
  "venv": ".venv"
}
```

## Use environment

On Mac, `source .venv/bin/activate` to activate the virtual environment.

The prompt should change to something like `(.venv) > virtenv: `.

## See what's installed

Use `pip freeze` to see a list of what's installed in the virtual environment.

## Stop using virtual environment

Use `deactivate` to stop using the virtual environment.
