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
