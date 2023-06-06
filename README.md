This is url shortener for Ozon internship.

In_memory storage is default option. To use postgres add flag -Memory_type psql

How Identifiers (short url) is generated : 
  <br>If identifier and url are passed with request, then it's short identifier for this link.
  
  <br>Otherwise its generated 1) using uuid 2) converted to base63 (a-zA-Z0-9_). Done this to protect from enumeration, so original_url cant be identified based on short_url. 
