# Changelog

## v0.2

- Fixed various type-o's in documentation.
- Renamed `IsSynthaxError` to `IsSyntaxError`.
- Fixed type-o in syntax error message.

## v0.1

 - **Ignore goveralls errors in travis build** (#c57cb74) by *Thomas de Zeeuw*, on *02 Apr 2016 12:31:57 UTC*. <br/>
*The service is to flaky to fail the build for*.
 - **Simplify test code** (#5257b94) by *Thomas de Zeeuw*, on *02 Apr 2016 11:19:24 UTC*. <br/>
*Found using https://github.com/dominikh/go-simple*.
 - **Remove unnecessary type conversions** (#88d4b5f) by *Thomas de Zeeuw*, on *15 Mar 2016 19:12:29 UTC*. <br/>
*Found using https://github.com/mdempsky/unconvert*.
 - **Fix two type-o's** (#57a5d6e) by *Thomas de Zeeuw*, on *09 Feb 2016 18:27:14 UTC*. <br/>
*Found with Misspell*.
 - **Add MIT license badge** (#c17a421) by *Thomas de Zeeuw*, on *09 Feb 2016 18:26:22 UTC*.
 - **Add Go Report Card** (#442052f) by *Thomas de Zeeuw*, on *09 Feb 2016 18:26:13 UTC*.
 - **Use shields.io badge for travis** (#d0acfbc) by *Thomas de Zeeuw*, on *09 Feb 2016 18:26:02 UTC*. <br/>
*Not as pixelated, https://shields.io*.
 - **Remove todo** (#045541f) by *Thomas de Zeeuw*, on *26 Jan 2016 15:25:10 UTC*. <br/>
*These have been converted into issues on GitHub*.
 - **Happy 2016** (#53efc69) by *Thomas de Zeeuw*, on *26 Jan 2016 11:51:38 UTC*.
 - **Make link to examples directory in ready** (#1f522b4) by *Thomas de Zeeuw*, on *10 Dec 2015 11:59:23 UTC*. <br/>
*Closes #17*.
 - **Display full utf-8 character in error messages** (#df315a8) by *Thomas de Zeeuw*, on *08 Dec 2015 20:27:17 UTC*. <br/>
*Closes #14*.
 - **Support simple custom types** (#b15cb7e) by *Thomas de Zeeuw*, on *07 Dec 2015 19:18:21 UTC*. <br/>
*Instead of getting the type of a reflect value get the kind (the go type) and use that in determining how to set the value.  Updates #16*.
 - **Allow # as a start of a comment** (#a98383c) by *Thomas de Zeeuw*, on *07 Dec 2015 19:00:25 UTC*. <br/>
*Closes #13*.
 - **Fix a lot of type-o's** (#6c1de5b) by *Thomas de Zeeuw*, on *07 Dec 2015 15:51:21 UTC*. <br/>
*seperator -> separator. qoute -> quote.  Closes #12*.
 - **Don't trim whitespace for quoted values** (#872a7a4) by *Thomas de Zeeuw*, on *07 Dec 2015 15:40:55 UTC*. <br/>
*Closed #11*.
 - **Update getNameParts doc** (#474ac87) by *Thomas de Zeeuw*, on *07 Dec 2015 15:39:23 UTC*. <br/>
*Type-o and start the line with a capital*.
 - **Small code cleanup** (#1327c25) by *Thomas de Zeeuw*, on *06 Dec 2015 15:52:37 UTC*. <br/>
*Move getConfigSectionsAlpha after first use*.
 - **Doc getSectionKeysAlpha** (#bcb8ca1) by *Thomas de Zeeuw*, on *06 Dec 2015 15:52:25 UTC*.
 - **Update examples in ready** (#759bdd9) by *Thomas de Zeeuw*, on *06 Dec 2015 15:37:46 UTC*.
 - **Create new Section type** (#8c9c321) by *Thomas de Zeeuw*, on *06 Dec 2015 15:34:23 UTC*. <br/>
*Closes #6*.
 - **Doc tags in Config.Decode** (#212ad3d) by *Thomas de Zeeuw*, on *06 Dec 2015 15:29:47 UTC*. <br/>
*Updates #4*.
 - **Use strconv.ParseBool for converting bools** (#8049e5e) by *Thomas de Zeeuw*, on *06 Dec 2015 15:22:53 UTC*.
 - **Add examples for using tags** (#3e79ab0) by *Thomas de Zeeuw*, on *06 Dec 2015 15:13:25 UTC*. <br/>
*Closes #7.  Updates #4*.
 - **Add decode example** (#1d6e07b) by *Thomas de Zeeuw*, on *06 Dec 2015 15:06:11 UTC*. <br/>
*Updates #7*.
 - **Fix simple log** (#a39522b) by *Thomas de Zeeuw*, on *06 Dec 2015 15:05:55 UTC*.
 - **Fix type-o in license info in source files** (#f6a6d90) by *Thomas de Zeeuw*, on *06 Dec 2015 14:57:47 UTC*.
 - **Create new examples** (#a7dcb87) by *Thomas de Zeeuw*, on *06 Dec 2015 14:54:58 UTC*. <br/>
*Starting with a simple example in the new `_examples` directory.  Updates #7*.
 - **Always check errors** (#a3d4201) by *Thomas de Zeeuw*, on *04 Dec 2015 18:42:39 UTC*. <br/>
*Even in examples*.
 - **Small doc update** (#4746a38) by *Thomas de Zeeuw*, on *04 Dec 2015 18:03:50 UTC*.
 - **Simplify parseSection code** (#ba51dac) by *Thomas de Zeeuw*, on *04 Dec 2015 18:03:35 UTC*.
 - **Remove TestPossibleQoute** (#3c72634) by *Thomas de Zeeuw*, on *21 Nov 2015 19:45:48 UTC*.
 - **Always quote keys and values in Config.String/Bytes** (#9f43ad8) by *Thomas de Zeeuw*, on *21 Nov 2015 19:44:42 UTC*.
 - **Split TestParse tests** (#07a5b27) by *Thomas de Zeeuw*, on *21 Nov 2015 19:38:04 UTC*.
 - **Fix test error output** (#b0e3cce) by *Thomas de Zeeuw*, on *21 Nov 2015 17:50:06 UTC*.
 - **Small code cleanup** (#f5f96ab) by *Thomas de Zeeuw*, on *21 Nov 2015 17:43:57 UTC*.
 - **Add coverage badge to readme** (#8d832c9) by *Thomas de Zeeuw*, on *21 Nov 2015 17:01:36 UTC*.
 - **Clean travis install commands** (#d25f35f) by *Thomas de Zeeuw*, on *21 Nov 2015 17:01:10 UTC*. <br/>
*Drop `get ./…`, no dependencies. Comment out `github.com/fzipp/gocyclo`, no used, but it should be. Also don’t update any of them*.
 - **Fix data race with testing with coverage** (#3b7bb4a) by *Thomas de Zeeuw*, on *21 Nov 2015 16:55:52 UTC*. <br/>
*Covermode count creates data races, atomic fixes this*.
 - **Add coveralls.io** (#b6974ed) by *Thomas de Zeeuw*, on *21 Nov 2015 16:24:21 UTC*.
 - **Remove glint from Travis** (#687a142) by *Thomas de Zeeuw*, on *04 Nov 2015 14:59:39 UTC*. <br/>
*Golint never exits with status code 1, so we’ll never see if it has any comments*.
 - **Fix travis.yml file** (#b32db3d) by *Thomas de Zeeuw*, on *04 Nov 2015 13:43:30 UTC*. <br/>
*Have to use spaces, tab aren’t allowed*.
 - **Travis more advanced testing** (#4f3083f) by *Thomas de Zeeuw*, on *04 Nov 2015 13:40:49 UTC*. <br/>
*Run gofmt simplify, go vet, deadcode, golint and go test with the race detector enabled.  Todo enable gocyclo*.
 - **Travis add OS X as a os to run tests on** (#898c968) by *Thomas de Zeeuw*, on *04 Nov 2015 13:39:18 UTC*.
 - **Remove unused code** (#e45e98e) by *Thomas de Zeeuw*, on *04 Nov 2015 13:34:34 UTC*.
 - **Fix license type-o** (#a1050b0) by *Thomas de Zeeuw*, on *15 Oct 2015 20:36:04 UTC*.
 - **Move getMapsKeysAlpha after first use** (#63eaa02) by *Thomas de Zeeuw*, on *25 Sep 2015 11:37:32 UTC*.
 - **Support ini tag in decoding structs** (#41fdae4) by *Thomas de Zeeuw*, on *25 Sep 2015 11:36:08 UTC*.
 - **Small test code cleanup** (#e8f406b) by *Thomas de Zeeuw*, on *25 Sep 2015 11:32:43 UTC*.
 - **Add go1.5 to Travis** (#4c1a7a5) by *Thomas de Zeeuw*, on *08 Sep 2015 12:07:43 UTC*.
 - **Check for OverflowError in TestDecodeValueOverflowError** (#aa77555) by *Thomas de Zeeuw*, on *08 Sep 2015 12:07:35 UTC*.
 - **Small code cleanup** (#21b0661) by *Thomas de Zeeuw*, on *18 Aug 2015 16:29:25 UTC*.
 - **Test Config.String, Config.Bytes** (#c63554d) by *Thomas de Zeeuw*, on *22 Jul 2015 13:33:50 UTC*.
 - **Add RFC3339, RFC1123 and RFC822 to time decoding** (#7ff0c15) by *Thomas de Zeeuw*, on *22 Jul 2015 13:33:49 UTC*.
 - **Drop unused code** (#51fc57b) by *Thomas de Zeeuw*, on *22 Jul 2015 13:33:49 UTC*. <br/>
*The check was in place for future code changes, but it's pointless*.
 - **type-o** (#1d28d4d) by *Thomas de Zeeuw*, on *21 Jul 2015 22:11:18 UTC*.
 - **Document the victory over go-fuzz** (#f233e1f) by *Thomas de Zeeuw*, on *21 Jul 2015 22:05:59 UTC*.
 - **Code cleanup** (#942bbf3) by *Thomas de Zeeuw*, on *21 Jul 2015 22:05:38 UTC*.
 - **Return error with duplicate section names** (#39c2ca0) by *Thomas de Zeeuw*, on *21 Jul 2015 21:43:51 UTC*.
 - **Stable v1** (#edfe11c) by *Thomas de Zeeuw*, on *21 Jul 2015 11:41:17 UTC*.
 - **Rename Scan, Config.Scan and ScanInto to Decode** (#f6b2e02) by *Thomas de Zeeuw*, on *21 Jul 2015 11:41:16 UTC*. <br/>
*Now it's Decode, Config.Decode and DecodeValue*.
 - **APi -> API** (#74d3e3c) by *Thomas de Zeeuw*, on *20 Jul 2015 15:00:29 UTC*.
 - **Add godoc examples** (#2128d72) by *Thomas de Zeeuw*, on *19 Jul 2015 20:56:29 UTC*.
 - **Update readme to match examples change** (#1bf9d21) by *Thomas de Zeeuw*, on *19 Jul 2015 20:45:19 UTC*.
 - **Move and update example** (#5146c33) by *Thomas de Zeeuw*, on *19 Jul 2015 20:43:25 UTC*.
 - **Drop fuzzer file** (#fd99e82) by *Thomas de Zeeuw*, on *19 Jul 2015 20:34:59 UTC*.
 - **Drop lineType and detection function** (#3948445) by *Thomas de Zeeuw*, on *19 Jul 2015 14:42:18 UTC*.
 - **Fix single character line section** (#a6b4404) by *Thomas de Zeeuw*, on *18 Jul 2015 14:29:51 UTC*.
 - **Add small fuzz test** (#a6c1f63) by *Thomas de Zeeuw*, on *18 Jul 2015 14:28:29 UTC*.
 - **Add test for just only opening section bracket** (#884390d) by *Thomas de Zeeuw*, on *18 Jul 2015 14:28:14 UTC*. <br/>
*Found by go-fuzz*.
 - **Fix golint/go vet issues** (#49f2a57) by *Thomas de Zeeuw*, on *11 Jul 2015 20:30:19 UTC*.
 - **Don't export the different errors** (#a1efed0) by *Thomas de Zeeuw*, on *11 Jul 2015 20:29:58 UTC*. <br/>
*The Is*Error() functions provide enough information*.
 - **Add ScanInto function** (#974c517) by *Thomas de Zeeuw*, on *11 Jul 2015 20:13:56 UTC*. <br/>
*This function scans a single (string) value into a variable*.
 - **Typo and small doc improvement** (#289349b) by *Thomas de Zeeuw*, on *11 Jul 2015 20:13:21 UTC*.
 - **Change error message value to single quote** (#5c06b20) by *Thomas de Zeeuw*, on *11 Jul 2015 20:11:05 UTC*. <br/>
*Error messages used to have a double qoute around the value, now we use a single qoute. Old: "ini: can't convert \"500\" to type uint8, it overflows the type", New: "ini: can't convert '500' to type uint8, it overflows the type"*.
 - **Export CovertionError** (#aee7fb7) by *Thomas de Zeeuw*, on *13 Jun 2015 21:18:41 UTC*.
 - **Export OverflowError** (#6ddba5d) by *Thomas de Zeeuw*, on *13 Jun 2015 21:18:30 UTC*.
 - **Export Synthax error** (#b9cd37d) by *Thomas de Zeeuw*, on *13 Jun 2015 21:18:08 UTC*.
 - **Rewrite error tests** (#da8ea77) by *Thomas de Zeeuw*, on *13 Jun 2015 21:16:59 UTC*. <br/>
*This merges the Is*Error test with Create*Error tests*.
 - **Add IsOverflowError and IsCovertionError functions** (#5bae5ef) by *Thomas de Zeeuw*, on *11 Jun 2015 20:18:18 UTC*.
 - **Rewrite error creation** (#9f6aca1) by *Thomas de Zeeuw*, on *11 Jun 2015 20:17:41 UTC*. <br/>
*Create types for covertion and overflow errors*.
 - **Improve overflow error message** (#53ed091) by *Thomas de Zeeuw*, on *11 Jun 2015 19:56:07 UTC*. <br/>
*No need to put the type twice in the message*.
 - **Improve testing code** (#fdea8cb) by *Thomas de Zeeuw*, on *11 Jun 2015 19:55:13 UTC*.
 - **Return overflow errors while scanning** (#8933012) by *Thomas de Zeeuw*, on *11 Jun 2015 19:31:46 UTC*.
 - **Prefix errors with `ini`** (#0195914) by *Thomas de Zeeuw*, on *11 Jun 2015 19:31:16 UTC*. <br/>
*Synthax errors are prefixed with `ini`, now so are the overflow and convert errors*.
 - **Merge parsing tests** (#40fc191) by *Thomas de Zeeuw*, on *03 Jun 2015 23:35:06 UTC*. <br/>
*Merge TestSectionLine and TestKeyValueLine into TestParse*.
 - **Drop temp file usage in TestComplete** (#28873a5) by *Thomas de Zeeuw*, on *03 Jun 2015 23:20:21 UTC*. <br/>
*We're using a buffer now, same result but quicker*.
 - **Make sure the section name isn't empty** (#6aa8dfb) by *Thomas de Zeeuw*, on *03 Jun 2015 23:10:01 UTC*.
 - **Improve doc** (#e4dc760) by *Thomas de Zeeuw*, on *03 Jun 2015 23:09:12 UTC*.
 - **Covert lineType to uint8** (#f80c844) by *Thomas de Zeeuw*, on *03 Jun 2015 23:08:53 UTC*. <br/>
*It should saves a little bit of memory*.
 - **Drop currentLine from the parser** (#969dff7) by *Thomas de Zeeuw*, on *03 Jun 2015 23:08:15 UTC*. <br/>
*It's not used in the code anymore after dropping the line from the synthax error*.
 - **Rename Global constant** (#1982130) by *Thomas de Zeeuw*, on *03 Jun 2015 23:07:23 UTC*. <br/>
*Rename it from "SUPERGLOBAL" to an empty string*.
 - **Add license info to error.go and error_test.go** (#fd71273) by *Thomas de Zeeuw*, on *03 Jun 2015 23:05:55 UTC*.
 - **Drop line from synthax error** (#4e2dbae) by *Thomas de Zeeuw*, on *03 Jun 2015 22:38:44 UTC*. <br/>
*It's unnecessary for the line to be in the error message*.
 - **Use a global name instead of the Global const in errors** (#2c1c592) by *Thomas de Zeeuw*, on *03 Jun 2015 22:33:44 UTC*.
 - **Improve parse error testing** (#fa1a4f7) by *Thomas de Zeeuw*, on *03 Jun 2015 22:29:58 UTC*. <br/>
*Move all parser errors tests into a single test. This give a better view on the whole error message and check if it's valid in stead of just a part of the message.  This also checks if the returned error is a synthax error*.
 - **Test in parallel** (#d458e8f) by *Thomas de Zeeuw*, on *03 Jun 2015 22:11:11 UTC*. <br/>
*Run all tets in parallel. It improves performance of running the tests, but also makes sure it's safe to use the package in parallel*.
 - **Update example to latest api** (#ab177a5) by *Thomas de Zeeuw*, on *02 Jun 2015 14:47:44 UTC*.
 - **Move the README example to its own directory** (#ddbd494) by *Thomas de Zeeuw*, on *02 Jun 2015 14:44:15 UTC*.
 - **Reorder constants** (#56f8ef4) by *Thomas de Zeeuw*, on *02 Jun 2015 14:36:43 UTC*.
 - **Drop forgotten todo item** (#bb58bd2) by *Thomas de Zeeuw*, on *11 May 2015 11:48:29 UTC*.
 - **Code cleanup** (#785535c) by *Thomas de Zeeuw*, on *11 May 2015 11:43:41 UTC*.
 - **Improve doc** (#0348b0a) by *Thomas de Zeeuw*, on *11 May 2015 11:37:58 UTC*.
 - **Take pointer in Config.Scan** (#fc79647) by *Thomas de Zeeuw*, on *11 May 2015 11:37:48 UTC*. <br/>
*Should reduce memory usage*.
 - **Move timeFormats to scan.go** (#e913568) by *Thomas de Zeeuw*, on *11 May 2015 11:37:02 UTC*. <br/>
*The variable is only used in scan.go so it's better to keep it their*.
 - **Remove Load** (#07781d0) by *Thomas de Zeeuw*, on *11 May 2015 11:18:52 UTC*. <br/>
*The ini package should focus on parsing the ini format, not loading files*.
 - **Rewrite of Scan** (#e118f10) by *Thomas de Zeeuw*, on *11 May 2015 11:06:04 UTC*. <br/>
*This is the complete rewrite of the scan functionality, it's now much cleaner code*.
 - **Fix TestComplete** (#f8052cd) by *Thomas de Zeeuw*, on *11 May 2015 11:00:41 UTC*.
 - **Drop pointless tests** (#d5a5b56) by *Thomas de Zeeuw*, on *11 May 2015 11:00:22 UTC*. <br/>
*TestLoadErrors and TestConfigString pointless, the tests are redundant*.
 - **Change Scan to take a io.Reader instead of a path** (#54db3ed) by *Thomas de Zeeuw*, on *11 May 2015 10:59:13 UTC*. <br/>
*The ini package should focus on parsing ini format, not on loading files*.
 - **Code cleanup** (#94b1c9e) by *Thomas de Zeeuw*, on *11 May 2015 10:57:30 UTC*. <br/>
*Move Config.WriteTo() above Config.buffer() In Config.Scan change the error message, no need to specifiy ini.Config.Scan, Config.Scan will suffice*.
 - **Fix error Config.String() not escaping qoutes** (#9918df9) by *Thomas de Zeeuw*, on *11 May 2015 10:51:45 UTC*. <br/>
*If qoute where used in a value or key Config.String() would not escape them, this change fixes that*.
 - **Add scan errors** (#7bfd666) by *Thomas de Zeeuw*, on *11 May 2015 10:47:25 UTC*. <br/>
*This commit adds createOverflowError and createConvertError functions*.
 - **Reorder functions** (#573dec4) by *Thomas de Zeeuw*, on *09 May 2015 14:37:57 UTC*. <br/>
*Reorder to match the switch statement in setReflectValue*.
 - **Better type detection in setReflectValue** (#98ef101) by *Thomas de Zeeuw*, on *09 May 2015 14:34:58 UTC*.
 - **Rewrite setReflectValue** (#65a35b1) by *Thomas de Zeeuw*, on *09 May 2015 11:16:37 UTC*. <br/>
*This rewrite splits up the function into multiple functions. It also improves error messages*.
 - **Rewrite Config.Scan** (#939b9e4) by *Thomas de Zeeuw*, on *09 May 2015 09:57:07 UTC*.
 - **Doc an code cleanup** (#f186558) by *Thomas de Zeeuw*, on *08 May 2015 23:07:58 UTC*.
 - **Add test for getConfigSectionsAlpha** (#1672a56) by *Thomas de Zeeuw*, on *08 May 2015 22:56:47 UTC*.
 - **Code cleanup** (#707411e) by *Thomas de Zeeuw*, on *08 May 2015 22:56:19 UTC*.
 - **Improve docs** (#15368b5) by *Thomas de Zeeuw*, on *08 May 2015 22:47:17 UTC*.
 - **Add Config.Bytes and Config.WriteTo methods** (#df085a5) by *Thomas de Zeeuw*, on *08 May 2015 22:42:37 UTC*.
 - **Rewrite Config.String()** (#a031104) by *Thomas de Zeeuw*, on *08 May 2015 22:41:42 UTC*.
 - **Update parse test** (#f8b7a9b) by *Thomas de Zeeuw*, on *08 May 2015 21:38:43 UTC*. <br/>
*Update with some test data and use reflecht.DeepEqual for testing of the configuration*.
 - **Add parse io error test** (#1974235) by *Thomas de Zeeuw*, on *08 May 2015 21:27:39 UTC*.
 - **Rewritten the parser** (#6d81b8f) by *Thomas de Zeeuw*, on *07 May 2015 20:39:21 UTC*. <br/>
*Completely rewritten the parser and put it in a new file*.
 - **Add synthax error** (#1d22bab) by *Thomas de Zeeuw*, on *07 May 2015 19:35:45 UTC*. <br/>
*This adds a new public function `IsSynthaxError(error) bool`, this can be used to determine if an error is synthax error (as the name suggests)*.
 - **Add copyright notice to all source files** (#bb573d3) by *Thomas de Zeeuw*, on *19 Apr 2015 22:22:01 UTC*.
 - **Fix lowercase section not being scanned** (#7c6f306) by *Thomas de Zeeuw*, on *19 Apr 2015 22:02:12 UTC*. <br/>
*Now "section name" resolves to "SectionName" while scanning*.
 - **Ignore extra sections in the Config corrected** (#882491c) by *Thomas de Zeeuw*, on *19 Apr 2015 21:59:50 UTC*. <br/>
*Previously the scanner would return an error if the Config had a section and the destination didn't, or if it's wasn't a struct.  Now it ignores the section*.
 - **Revert "Ignore extra sections in the Config"** (#e6bc058) by *Thomas de Zeeuw*, on *19 Apr 2015 21:45:49 UTC*. <br/>
*This reverts commit 637828b1030e53477ee125c91602099ddb14924f*.
 - **Ignore extra sections in the Config** (#637828b) by *Thomas de Zeeuw*, on *19 Apr 2015 21:42:21 UTC*. <br/>
*Previously the scanner would return an error if the Config had a section and the destination didn't, or if it's wasn't a struct.  Now it ignores the section*.
 - **Add yyyy-mm-dd to scanning of time** (#6a8f378) by *Thomas de Zeeuw*, on *11 Apr 2015 21:14:11 UTC*.
 - **Load return possible bufio.Scanner error** (#5c051c1) by *Thomas de Zeeuw*, on *10 Apr 2015 19:33:08 UTC*.
 - **Fix int overflow checking in Scan of slices** (#382d614) by *Thomas de Zeeuw*, on *06 Apr 2015 16:31:54 UTC*.
 - **Check for int overflow in Scanning of slices** (#48d18d1) by *Thomas de Zeeuw*, on *06 Apr 2015 16:29:04 UTC*.
 - **Clean up todos** (#c408f7b) by *Thomas de Zeeuw*, on *06 Apr 2015 16:28:43 UTC*.
 - **Don't set keys on maps** (#5aa64bc) by *Thomas de Zeeuw*, on *06 Apr 2015 16:28:17 UTC*.
 - **Improve readme example** (#170b08f) by *Thomas de Zeeuw*, on *05 Apr 2015 20:27:11 UTC*.
 - **Init()** (#49b4078) by *Thomas de Zeeuw*, on *05 Apr 2015 20:20:50 UTC*.
