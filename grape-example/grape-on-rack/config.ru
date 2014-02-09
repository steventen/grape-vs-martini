require 'grape'
require 'grape-entity'
require 'active_record'
Dir[File.dirname(__FILE__) + '/*.rb'].each {|f| require f}

ActiveRecord::Base.establish_connection(
  :adapter => 'mysql2',
  :database => 'grape_vs_martini_api',
  :host => "127.0.0.1",
  :pool => 50,
  :username => "root",
  :password => "abc123"
)

#\ -s puma
run MySite::API