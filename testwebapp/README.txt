
  gaerecordsapp
  
  Test app to show a Google App Engine app using gaerecords to persist records
  in the datastore.
  
  ---
  
    This project uses:
    
      gaerecords        https://github.com/matryer/gae-records
      goweb             RESTful web framework - http://goweb.googlecode.com/
      mustache.go       Web templating library - https://github.com/hoisie/mustache.go
    
  ---
  
    Run the local webserver:
    
      do this in Terminal:
      
        /Users/matryer/Work/lib/google_appengine/dev_appserver.py /Users/matryer/Work/gae-records/testwebapp
      
      (NOTE: You'll have to modify these paths to suit your username)
    
    Then Visit: 
  
      http://localhost:8080/
    
  ---
  
    Fix errors:
    
      1. ERROR: can't find import: goweb
      
      - Get goweb from http://code.google.com/p/goweb/source/checkout
      - cd to the goweb dir (usually, goweb/goweb)
      - gomake install
      - copy goweb.a from the _obj folder
        and paste it into google_appengine/goroot/pkg/*/goweb.a
        
        
      2. "Oops, something went wrong: No route found for that path"
      
        visit http://localhost:8080/people