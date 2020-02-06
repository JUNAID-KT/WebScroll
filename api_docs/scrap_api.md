**Fetch content of a given URL**
----
    
* **Description**
    
       This API will fetch web content based on the given URL

* **Version**

       /v1.0
  
* **URL**

      localhost:8080/v1.0/web_scroll/scrap
       
* **Method:**

       POST
  
* **URL Params**

        None

* **Data Params**

      {
         "url": "google.com",
      }

* **Success Response:**

  * **Code:** 200 OK 
    
    **Content:** 

        {
             "status_code": 200,
             "description_code": "OK",
             "description": "Web content saved"
        }
 
* **Error Response:**

  * **Code:** 500 INTERNAL SERVER ERROR 
   
    **Content:** 
        
         {
            "status": {
            "status_code": 500,
            "description_code": "FAILURE",
            "description": "Post http://localhost:9200/ethereum_transactions/_search: dial tcp 127.0.0.1:9200: connect: connection refused"
            }
        }

  OR

  * **Code:** 400 STATUS BAD REQUEST
               
    **Content:** 
    
        {
          "status": {
            "status_code": 400,
            "description_code": "FAILURE",
            "description": "Request binding failed"
          }
        }

