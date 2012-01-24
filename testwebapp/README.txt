
  gaerecordsapp
  
  - test app to show a Google App Engine using gaerecords to persist records
    in the datastore.
    
    
  ---
  
    Run the local webserver:
    
      /Users/matryer/Work/lib/google_appengine/dev_appserver.py /Users/matryer/Work/gae-records/testwebapp
    
    Visit: 
  
      http://localhost:8080/
    
  ---
  
    Fix errors:
    
      ERROR: can't find import: goweb
      
      - Get goweb from http://code.google.com/p/goweb/source/checkout
      - cd to the goweb dir
      - gomake install
      - copy goweb.a from the _obj folder
        and paste it into google_appengine/goroot/pkg/*/goweb.a
        
