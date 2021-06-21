.. _alfred-reference:

Alfred
######

Alfred is the command-line interface included with Go-Web. It provides a number of helpful commands that can assist you while you build your application.
You can compile `Alfred` by running `sudo make build` in your project root.

Usage:    `./alfred --help`

.. list-table:: Alfred available commands
    :widths: 50 50
    :header-rows: 1

    * - Commands
      - Descriptions
    * - --help, -h
      - Shows help menu
    * - --make-controller, -mC <controller_name>
      - Creates new Go-Web controller*
    * - --make-command, -mCMD <command_name>
      - Creates new Go-Web command*
    * - --make-model, -mM <model_name>
      - Creates new Go-Web model*
    * - --make-migration, -mMDB <migration_name>
      - Creates new Go-Web SQL migration*
    * - --make-middleware, -mMW <middleware_name>
      - Creates new Go-Web SQL middleware*
    * - --make-job, -mJ <job_name>
      - Creates new Go-Web async job*
    * - --show-route, -sR
      - Shows service routes*
    * - --migrate-up, -mU
      - Executes migrations*
    * - --migrate-rollback, -mR <steps>
      - Executes migrations rollback*
    * - --generate-key, -gK
      - Generate new application key*

\* *Run this command only in project root*