# francocorrea.com 

[![Greenkeeper badge](https://badges.greenkeeper.io/francocorreasosa/francocorreacom.svg)](https://greenkeeper.io/)

This is my personal website/blog.

## Database Setup

This website stores its data in a PostgreSQL database. You can set it up by looking at the `database.yml` file.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now the tool can create the databases in that file for you:

	$ buffalo db create -a

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see the site.


[Powered by Buffalo](http://gobuffalo.io)
