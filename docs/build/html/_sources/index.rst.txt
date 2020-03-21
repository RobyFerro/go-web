Go-Web framework documentation
==================================

Why Go-Web?
-----------
Go-Web adopts a “convention over configuration” approach similarly to frameworks like Laravel ,Symfony and Rails.

By following this principle, Go-Web reduces the need for “repetitive” tasks like explicitly binding
routes, actions and middleware; while not problematic per se, programmers may want to adopt
“ready-to-use” solutions, for instances if they want easily build complex services or have a non-Go
background, aiming at a reduced “productive lag” which could be experienced when approachingnew technology.

Programmers may want to use existing frameworks like Gin-Gonic and Go Buffalo, but
Go-Web differs from them because of the aforementioned “convention over configuration” approach.

.. toctree::
   :maxdepth: 2
   :caption: Table of Contents

   installation
   architecture
   dev_environment
   service_configuration
   cli
   controller
   routing
   run
   middleware
   async
   database
   authentication
   front-end