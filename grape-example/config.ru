require 'active_record'
Dir[File.dirname(__FILE__) + '/*.rb'].each {|f| require f}

#\ -s puma
run MySite::API