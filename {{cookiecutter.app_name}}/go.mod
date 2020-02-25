module {% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}

require (
	github.com/sirupsen/logrus v1.4.2
)
