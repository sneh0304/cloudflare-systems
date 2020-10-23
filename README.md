# cloudflare-systems
To build this program, run 'make build' in the terminal. It will geenrate a executable file cloudflare_systems.\
To run the program, run the executable file using './cloudflare_systems --url your_url --profile #times you want to hit the url'

# commands
make build\
./cloudflare_systems --url https://cloudflare-test.snehashish.workers.dev/links --profile 10

# Results
HTTP/1.1 200 OK\
Date: Thu, 22 Oct 2020 23:58:37 GMT\
Content-Type: application/json\
Content-Length: 332\
Connection: close\
Set-Cookie: __cfduid=d6f2cfd603c3308befc8c14a3851867381603411117; expires=Sat, 21-Nov-20 23:58:37 GMT; path=/; domain=.snehashish.workers.dev; HttpOnly; SameSite=Lax\
cf-request-id: 05f4591e6d000074418d21e000000001\
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"\
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report?lkg-colo=16&lkg-time=1603411118"}],"group":"cf-nel","max_age":604800}\
NEL: {"report_to":"cf-nel","max_age":604800}\
Server: cloudflare\
CF-RAY: 5e672add7e7e7441-IAD

[{"name":"IoT based weather station","url":"https://ieeexplore.ieee.org/document/7988038/"},{"name":"Flow based environmental monitoring for smart cities","url":"https://ieeexplore.ieee.org/document/8125882"},{"name":"9gag","url":"https://9gag.com/"},{"name":"Cloudflare&nbsp;Workers","url":"https://developers.cloudflare.com/workers/"}]

# Results for www.amazon.com
<img src="https://github.com/sneh0304/cloudflare-systems/blob/main/amazon.png"/>

# Results for www.google.com
<img src="https://github.com/sneh0304/cloudflare-systems/blob/main/google.png"/>

# Results for www.cloudflare.com
<img src="https://github.com/sneh0304/cloudflare-systems/blob/main/cloudflare.png"/>

# Results for general_assignment
<img src="https://github.com/sneh0304/cloudflare-systems/blob/main/cloudflare-test.snehashish.workers.dev.png"/>

# Results for links
<img src="https://github.com/sneh0304/cloudflare-systems/blob/main/links.png"/>
