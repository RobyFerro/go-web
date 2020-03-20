Asynchronous jobs
=================
Go-web allows developers to create and schedule asynchronous jobs that will be dispatched in a queue. Like controllers, models and other entities, a job can be created with the CLI by running command:

.. highlights:: ./goweb job:create <job name>

Go-Web uses Redis to manage queues and developers can handle jobs with functions Schedule and Execute.

The following listing illustrates an example of a Go-Web job:

.. code-block:: go

    data := job.MailStruct {
        Message: "Hello world",
        To: []string { "test@test.com", "test@test1.com" },
    }

    payload, _ := json.Marshal(data) j := job.Job {
        Name: "Send email"
        MethodName: "Mail"
        Params: job.Param { Name: "message", Payload: string(payload), Type: "int" },
    }

    j.Schedule("default", c.Redis)

Once scheduled, a job can be run with CLI command:

.. highlights:: ./goweb queue:run <queue name>

The default queue can be run with command:

.. highlights:: ./goweb queue:run default