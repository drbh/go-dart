CURL *hnd = curl_easy_init();

curl_easy_setopt(hnd, CURLOPT_CUSTOMREQUEST, "GET");
curl_easy_setopt(hnd, CURLOPT_URL, "http://localhost:3000/api/students");

curl_easy_setopt(hnd, CURLOPT_COOKIE, "server-session-cookie-id-for-alice_wallet=s%253A7n3HoAlDTEjeUOehLrhVfICrqxIKWIh1.zsRShRgxck%252BWYuKnC5gW37dkoRcA5EbiJXwrnUsqWr8");

CURLcode ret = curl_easy_perform(hnd);