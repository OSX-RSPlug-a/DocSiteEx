# DocSiteEx
Project of documentation site

Dependencies:

    - Gofiber;
    - Mkdocs;
    - Material for MkDocs;


Install dependencies:

      go get github.com/gofiber/fiber
      
      pip install mkdocs
  
      pip install mkdocs-material


Init the project:

      mkdocs new my-project
      
      mkdocs build

      go run  main.go

You can initialize also with:

      go build main.go
  
      chmod +x main            // this main It's a binary that was compiled
  
      nohup ./main </dev/null >nohup.out 2>nohup.err &      // to execute in background the script or program 
  
