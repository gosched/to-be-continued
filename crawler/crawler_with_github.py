import getpass
import requests
from pyquery import PyQuery as pq

class Login(object):
    def __init__(self):
        base_url = 'https://github.com/'
        self.login_url = base_url +'login'
        self.post_url = base_url +'session'
        self.logined_profile_url = base_url +'settings/profile'
        self.session = requests.Session()
        self.session.headers = {
            'Referer': 'https://github.com/',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36',
            'Host': 'github.com'
        }

    def get_token(self):
        response = self.session.get(self.login_url)
        doc = pq(response.text)
        token = doc('input[name="authenticity_token"]').attr("value").strip()
        return token
    
    def request_login(self, email_or_username, password):
        token = self.get_token()
        post_data = {
            'authenticity_token': token,
            'login': email_or_username,
            'password': password,
        }

        response = self.session.post(self.post_url, data=post_data)        
        # response 302 redirect to 'https://github.com/'
        print()
        print(f"response url：{response.url}")
        if response.status_code != 200:
            return
        self.check_homepage(response.text)
        self.check_profile()

    def check_homepage(self, html):
        doc = pq(html)
        user_name = doc("summary > span").text().strip()
        print(f"username：{user_name}")

        repositories = doc("div.Box-body > ul > li").text().split()
        for repository in repositories:
            print(repository)
    
    def check_profile(self):
        response = self.session.get(self.logined_profile_url)
        if response.status_code != 200:
            return
        html = response.text
        doc = pq(html)
        title = doc("title").text()
        name = doc("#user_profile_name").attr("value")
        bio = doc("#user_profile_bio").text()
        location = doc("#user_profile_location").attr("value")
        url = doc("#user_profile_blog").attr("value")
        print(f"Page title: {title}")
        print(f"Name: {name}")
        print(f"Bio: {bio}")
        print(f"URL: {url}")
        print(f"Location: {location}")

    def try_login(self):
        email_or_username = input("Please input email or username: ")
        password = getpass.getpass("Please input password: ")
        self.request_login(email_or_username=email_or_username, password=password)

if __name__ == "__main__":
    login = Login()
    login.try_login()

# from urllib import request, response, parse, error
# import ssl

# def login(target):
#     print('Login to github.com')
#     try:
#         username_or_email = input('Username or Email: ')
#         passwd = input('Password: ')

        # login_data = parse.urlencode([
        #     ('login', username_or_email),
        #     ('password', passwd),
        # ])

        # client_request = request.Request('https://github.com/login')
        # client_request.add_header('Origin', 'https://github.com/')
        # client_request.add_header('User-Agent', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36')

        # context = ssl._create_unverified_context()
        
        # with request.urlopen(client_request, context=context, data=login_data.encode('utf-8')) as f:
            # print(type(f))
            # print(dir(f))
            # print('Status:', f.status, f.reason)
            # data = f.read()
            # print('Headers:')
            # for k, v in f.getheaders():
            #     print(f"[{k}] -> [{v}]")
                # print('%s: %s' % (k, v))
            # print(type(data))
            # print(dir(data))
            # print('Data:', data.decode('utf-8'))

    # except (error.URLError, IOError) as e:
    #     print(e)

# target_urls = [
#     '1https://github.com/',
#     'https://github.com/',
# ]

# for url in target_urls:
#     login(url)