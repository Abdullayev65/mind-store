Mind storage <a name="TOP"></a>
===================

- - - - 

# What does it do? #

    You can store your thoughts. 
    It similar with Notion however 
    here you can search for others 
    thoughts about what you are interesting 
    in and you soting your thoughts 
    with hashing(then only you can read 
    what you writed, even admins can not see 
    what you writed case it stores with hashing).

## What technologies were used ##

    Golang - programming language
    Gin - router
    Sqlx - database primary tool
    hashid - tool for hashing ids
    defoultcase - helping json case [like FirstName to "first_name"]
    
    main tool are these. 
    src down below

hashid links [github](https://github.com/speps/go-hashids), [youtube](https://www.youtube.com/watch?v=tSuwe7FowzE);

## Default case ##

    it helps don't write same logic every time
    
    usually we write like this:
    type User struct {
        Id        int    `json:"id"`
	    FirstName string `json:"first_name"`
    }

    whith this it works like this:
    defaultcase.SetDefaultCase(defaultcase.Snak_case)
    type User struct {
        Id        int
	    FirstName string
    }
    
    same json is generated in both above cases:
    {
    "id":0,
    "first_name":""
    }

front-end written in vue.js [github](https://github.com/Abdullayev65/mind-store-front);
