##Errors package provides simple way to handle application and API errors 

####Types
There are following error categories which allows to distinguish them from each other
 * PublicError - represents custom generic error which should be exposed to end user
 * PrivateError - represents internal errors which must not be shown to end user
 * UnmarshalError - represents errors of incorrect body format
 * ValidationErrors - represents errors formed by validator.v8 package 

Each error type implements TypedError interface.

###How to use it?
1. Add middleware to your router.
2. Create an error object from one of described above types.
3. Call errors.AddErrors(context, yourObject).

####Middleware
The package comes with handy gin middleware. The middleware processes incoming errors.
Simply call `errors.ErrorHandler(logger)` to get the middleware, then it can be added to your router 
middleware stack.

####Error details
##### Public Error
PublicError needs only Code and HttpStatus. Other fields are optional

##### Private Error
PrivateError needs Message will be passed to logger.Error (that used in middleware) method.

You also can add custom log pairs with AddLogPair method.

##### Unmarshal Error
UnmarshalError needs only pointer to json.UnmarshalTypeError. Usually this error is created by gin and converted by helper methods.

##### Validation Errors
This type of error has list of validation errors.

Code and Source are mandatory fields.

You can use it to create custom validation error or convert validator error with helper methods.


#### Helpers
Method ShouldBindToTyped should be used to convert gin error to TypedError and then pass it to context with errors.AddErrors. You can also do it with one method AddShouldBindError.

This method converts error to UnmarshalError or ValidationErrors.

Example:
```go
    if err := context.ShouldBindJSON(&myform); err != nil {
    	errors.AddShouldBindError(context, err)
    	return
    }
```



### Formatters

Formatters are used to convert validation error specified by tag to default error format.

Package has list of default formatters for common validations like max, min, length, etc. All them you can find in error_formatters.go file.

##### Defining own formatters

To define you own formatters you need to build map of [tag => ValidationErrorFormatter] and call SetFormatters in the app initializing.

Package will search formatter by tag in new formatters then in default. This allows to redefine existing formatters.

### Versions

0.1.0 version uses v8 validator.  
1.0.0 version uses v10 validator.