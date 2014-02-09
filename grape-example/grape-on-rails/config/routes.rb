GrapeApiRails::Application.routes.draw do
  require 'api'
  mount GrapeApiRails::API => "/"
end
