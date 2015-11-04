# jive-sdk-go
![Alt](/dev_logo.png "Jive Developer Logo")

This is currently a placeholder for a series tools to assist Go developers to build Jive Add-Ons. Reach out in the [Jive Developer Community](community.jivesoftware.com/community/developer) if you are interested in contributing.

**This SDK is currently a work in progress.** :smile:

Please reference the [Jive-SDK-Your-Way Bounty](https://community.jivesoftware.com/community/developer/blog/2015/11/03/jive-sdk-your-way-contribute-an-sdk-and-get-paid).

# To-do
- [ ] Validate headers -- register and unregister [Reference](https://community.jivesoftware.com/docs/DOC-99941#jive_content_id_Ensure_Register_Calls_are_Coming_from_a_Genuine_Jive_Instance)
- [ ] Validate signed fetch add-on requests [Reference](https://community.jivesoftware.com/docs/DOC-156557)
- [ ] Examples Framework [Reference](https://community.jivesoftware.com/docs/DOC-99941)
  - [ ] Tile (with **/jive/tile/register** and **/jive/tile/unregister** endpoints)
  - [ ] App (with static assets -- html, css, js)
  - [ ] Add-On (with configure screen and **/jive/addon/register** and **/jive/tile/unregister** endpoints)
- [ ] Documentation (include method defintions in README.md)

#Testing Validation
**Sample Payload**:
{ clientId: '96pndr7f9t6k7kt10lpmpg73w0krikvf.i',
  code: 'i38k1q1w1dhheijudu2n9riaeqlsno2q.c',
  scope: 'uri:/api',
  tenantId: 'b22e3911-28ef-480c-ae3b-ca791ba86952',
  jiveSignatureURL: 'https://market.apps.jivesoftware.com/appsmarket/services/rest/jive/instance/validation/8ce5c231-fab8-46b1-b8b2-fc65deccbb5d',
  clientSecret: 'nhi8a9mcpg51tdx7hwhlbxj5biutck.s',
  jiveSignature: '2ZNaB5hnn8gJgSELneU//6D2T0+n/rUf9uIFDKyRlAo=',
  jiveUrl: 'https://sandbox.jiveon.com',
  timestamp: '2015-11-04T00:39:25.361+0000' }

**Sample Hex Digest Output**: 
1b32eea102a83b3ea74ac550a770283ab35dc87033ab115c4a1d1a784d998ca6


# Begin Documentation
## method(param1, param2)
This is an example documentation with information about the parameter.

### Response Codes
Success : `200`
Failure : `403`