# jive-sdk-go
![Alt](/dev_logo.png "Jive Developer Logo")

This is currently a placeholder for a series tools to assist Go developers to build Jive Add-Ons. Reach out in the [Jive Developer Community](community.jivesoftware.com/community/developer) if you are interested in contributing.

**This SDK is currently a work in progress.  As I get time, I will continue to try and update this along with the other example languages. Feel free to contribute**:smile:

# To-do
- [ ] Routes
- [ ] Validate headers -- register and unregister
- [ ] Validate signed fetch add-on requests
- [ ] Example Add-on Framework
- [ ] Examples Included
  - [ ] Activity Stream Integration
  - [ ] Basic Jive App
  - [ ] Tile - Gallery
  - [ ] Tile - Carousel
  - [ ] Tile - List
  - [ ] Tile - Table
  - [ ] Tile - Custom View

# Routing
There are two endpoints the SDK supports **/register** and **/unregister**

##/register
Requested by Jive instance when an add-on is first installed or "reconnect service" is selected from Jive Add-on Manager. Up to 3 requests are made by Jive to connect to the endpoint.  

Methods: `POST`
Success: `200`

Error:
  `400 (Abort) - Not JSON`
  `401 (Fail) - Unable to register`

##/unregister
Requested by Jive instance when the add-on is uninstalled.
Methods: `POST`
Success: `200`
Error:
  `400 (Abort) - Not JSON`
  `401 (Fail) - Unable to register`

#Header Validation
There are two validations required, both using the SHA-256 hash implementation.

1. Logic to validate add-on registration request is originating from an authentic Jive Instance by checking with the Jive Marketplace URL
2. Verify that a signed-fetch request from an Opensocial container is valid.

More details about validating signed requests can be found in the following documents from the Jive Developer Community:
[https://community.jivesoftware.com/docs/DOC-156557](https://community.jivesoftware.com/docs/DOC-156557)
[https://community.jivesoftware.com/docs/DOC-99941](https://community.jivesoftware.com/docs/DOC-99941)

##(Un)Register Payload
Registration and unregistration of the add-on will call their respective endpoints. During registration, Jive will **POST** to the **/register** endpoint with a JSON object with the structure outlined below:

```javascript
{
	“clientID”  : “xxxxxxx”,
	“code” : “xxxxxxx”,
	“scope” : “xxxxxxx”,
	“tenantID” : “xxxxxxx”,
	“JiveSignatureURL” : “xxxxxxx”,
	“clientSecret” : “xxxxxxx”,
	“jiveSignature” : “xxxxxxx”,
	“jiveURL” : “xxxxxxx”,
	“timestamp” : “xxxxxxx”,
}
```

During an unregister of the add-on, the same payload is sent to the **/unregister** endpoint with the exception of the **clientSecret**. In order to validate the request, the SHA-256 digest of the clientSecret from the original registration request must be preserved.

*Note: If the service is shut down, the add-on will need to be registered from Jive’s Add-on Manager due to the clientSecret.*

##Validating (Un)Register Method
Constructing the register and unregister validation requires a few steps to ensure authenticity of the source. First we need to get the SHA-256 digest value of the Client Secret—we will be sending this to the URI value in the JiveSignatureURL key of the payload. We will also be removing the JiveSignature from the request JSON since that’s what we’re checking for and set it in a header as a value of the key X-Jive-MAC.
POST Add-on Validation (Un)Register Request
A typical CURL request to validate the source would be as follows:
```
curl –X POST –d ‘{“clientID”:”xxxx”,”code” :”xxxx”,”scope” :”xxxx”,”tenantID” :”xxxx”,”jiveSignatureURL” :”xxxx”,”clientSecret”: ”xxxx”,”jiveURL” :”xxxx”,”timestamp” :”xxxx”}’ –H ‘{“X-Jive-MAC” : <jiveSignature>}’  <jiveSignatureURL>
```

###Status Codes From Jive Marketplace URL
Success: `204`
Failure: `403`

##Validate Signed-Fetch Request from Opensocial Container
This method allows the Jive App/Tile to include the actor identity to be included in the request. The request will include an **Authorization** header that begins with **JiveEXTN**. This is an important indicator to knowing that the request is signed.

We then need to create a object of a substring of the Authorization header (EXCLUDING the string **JiveEXTN** and the **signature** key/value) and parse it into key value pairs.

The suffix (“**.s**”) from the **clientSecret** must be removed as well before we Base64 decode it.

##Expected Signature
To check the signature received in the header with what it should be, we must take the SHA-256 digest of clientSecret (Base64 decoded) and the object created with the rest of the **Authorization** header. This then can be Base64 encoded and remove the URL encoding to give the **expected signature**. The **signature** in the **Authorization** header should match the **expected signature**.

#Building an Example Add-on Framework
Jive Add-on’s are packaged into an archive file that can be installed in a Jive Community to expand and enrich the instance. Possible types of Add-on’s include apps, tiles, activity streams, external storage frameworks, and other Jive extension types. Examples of each type of add-on is encouraged and should be kept in their in their respective folders within the /examples/ folder of the language/framework’s SDK.

##Base Files & Folders
All Jive Add-on’s must contain a meta.json and definition.json file. All other files and configuration pages are optional based on the Add-on type and fields specified in the meta.json.
- definition.json – Schema specific to the add-on type
- meta.json – Contains several fields that are mostly optional but can be used to override defaults and also set the “type” field
- /data/ - Contains files related to the Add-on, including static HTML, JS, CSS, icons that will be served to the client
- /l18n/ - Java properties formatted resource bundles for localization support

##Exported File
After the user selects an add-on type, s/he will receive an archive file named **extension.zip** containing the related resources.

More details about building a Jive Add-on can be found in the Jive Developer Community:
[https://community.jivesoftware.com/docs/DOC-99941](https://community.jivesoftware.com/docs/DOC-99941)

