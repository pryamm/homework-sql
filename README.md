Build a minimal service to create, update, get, and delete movies, using this boilerplate.
- The system should receive an external GET request with movie slug, ex: titanic,
return 200 (ok) if found and 404 (not found) if not found
- The system should receive an external POST request with create movie API
contract, with status code 201 (created) if already successful to create. Movie slug
must be unique
- The system should receive an external PUT request with slug as movie identifier
alongside with update movie API contract and status code 200 (ok) if successfully
update movie content, 400 (bad request) if there is a client error, 422
(unprocessable entity) if there’s a duplicate movie slug
- The system should receive an external DELETE request with slug as movie identifier
return 200 (ok) if found and success to delete and 404 (not found) if the movie not
found
- Solution should be running
Bonus:
- It will be a huge plus if add unit tests
- Add status code accordingly based on the errors
Step:
- Feel free to use any library you prefer to handle the request, validation, mock etc.
- Create a private repository, invite your class tutor and rysmaadit as collaborator in
your github with the code, and a README.md with this text and any further
information to explain your solution
