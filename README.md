# NFactorial meetups tracker

This is an elective project given for me by NFactorial team. Let's see whether I pass!

## Building and running the project
1. Dependencies
  - Go version 1.22.2 or higher
  - GCC
  - (Optional) sqlite3 and sqlitebrowser, if you want to play with the project by adding your own tables
2. Build the project and run
Go into a root directory of a project and either:
```bash
go run ./cmd/web
```
or
```bash
go build ./cmd/web
./web # or ./web.exe if you're running on Windows

```
Afterwards go to `http://127.0.0.1:4000/` on your browser and you should see the application

## The development process
Let me share how the development of this project went, since this is a part of the requirements.
I had an experience with back-end development before, but not as much as in frontend and as such I wanted the most bare-bones
and hassle-free approach. I had experience with Django, NodeJS and Go, and decided to choose the latter to write the project.
The reason itself is very simple: the language itself is very straightforward in every process of development, starting with writing code
and ending with building it and deploying. No need to have a bunch of `node_modules` (Hello NodeJS), or install `python` with 
`python-venv` in order to build project in a way which does not make a mess out of your home folders. Everything is very simple,
with everything I need already being at hand in the Go standard library.

When I wrote applications using any business logic, most of the time I had all of it in the back-end. That's how most of big
frameworks like Django want you to work in general, but in this case I decided to use a different approach: instead of having all 
of the logic on the server side, which as it turns out is rather expensive for business, I decided to write a SPA, which does
most of its work through AJAX requests which are sent to a back-end API. That way both front-end and the back-end work together!

The project structure is very simple for the sake of getting things done (I also had to learn and revise Go that I did not write in for a while, so there's also a time constraint).
Thus the rpoject isn't really scalable: methods which interact with a database and the ones which control api are in the same file.
For the sake of convenience SQLite is shosen as a database, so that a reviewer could run it in no time, albeit real apps would use
something more feature-complete, like PostgreSQL or MySQL. The front-end does not use a framework of any kind, it's just basic javascript with dom elements hardcoded in strings. 
Yet for this project I consider it to be quite featureful: it has basic routing and renders everything on its own using AJAX calls to
query information.

There are some pros to using this kind of approach (i.e. building a SPA) when working on a project:
1. There is a clear separation of concerns between the front-end and the back-end, so this means that front-end developers 
can focus more on getting better user experience, while back-end developers can focus on building a more stable and robust API
2. No need to serve an entire HTML page for every user request or change, every chage is incremental which reduces the load on the server
3. No redirections, as such the appliation is faster as well

But there are some drawbacks to this as well:
1. Getting SEO optimixation is much harder than in traditional websites
2. Requires a lot of javascript code to be written, therefore a framework is preferred, compared to my vanilla JS approach
