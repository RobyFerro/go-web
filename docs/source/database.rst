Database
========

Models
------
In MVC frameworks models are responsible of the database interaction logic. Go-Web take advantage of GORM library to provide them.

To create a new model you can use its specific CLI command:

.. highlights:: ./goweb model:create <model name>

Models are located in: *<go-web>/database/model*

Migration
---------
Migrations are like version control for your database, allowing your team to easily modify and share the application’s database schema.
Developers can creates new migration as follows:

.. highlights:: ./goweb migration:create <migration_name>

Developer can find its newly created migration files in <go-web>/database/migration directory.

Seeding
-------
By implementing "Seed" method you'te able to seed your table with a "fake" data.
Go-Web uses *https://github.com/brianvoe/gofakeit* as faker library.

See the below example:

.. code-block:: go

    // Execute model seeding
    func (User) Seed(db *gorm.DB) {
	    for i := 0; i < 10; i++ {
		    password := gofakeit.Password(true, true, true, true, false, 32)
		    encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

		    user := User{
			    Name:     gofakeit.FirstName(),
			    Surname:  gofakeit.LastName(),
			    Username: gofakeit.Username(),
			    Password: string(encryptedPassword),
		    }

		    if err := db.Create(&user).Error; err != nil {
			    exception.ProcessError(err)
		    }
	    }
    }


