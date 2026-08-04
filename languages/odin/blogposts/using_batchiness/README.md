# Batchiness Tutorial Source Code

Source code for the [Using Batchiness](https://nofuss.co.za/blog/using_batchiness/) blog post.

## Suggested Directory Structure

```
project_root_dir/
├── vendor/
│   └── batchiness/
│       ├── batch/          ← the Odin package you import
│       └── web/
│           └── batchiness.js  ← JS glue you copy/link to your web dir
├── src/
│   └── main.odin
├── web/
│   ├── index.html
│   └── batchiness.js      ← copy from vendor/batchiness/web/
└── build.sh
```
