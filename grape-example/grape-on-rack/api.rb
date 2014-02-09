module MySite
	class Project < ActiveRecord::Base
	  belongs_to :company
	end

	class Company < ActiveRecord::Base
	  has_many :projects
	end

	class API < Grape::API
		format :json
		default_format :json
		# prefix 'api'
		cascade false
		default_error_formatter :json

		helpers do
	    def current_company
	      key = params[:key]
	      @current_company ||= Company.where(:api => key).first
	    end

	    def authenticate!
	      error!({ "status" => "Fail", "error_message" => "Bad Key" }, 401) unless current_company
	    end
	  end

		rescue_from :all do |e|
	    Rack::Response.new({ "status" => "Fail", "error_message" => e.message }.to_json, 405)
	  end

		before do 
			authenticate!
		end

		get "projects" do
			projects = current_company.projects
			present :data, projects, :with => APIEntities::Project
			present :status, "Success"
		end

		get "projects/:id" do
			project = current_company.projects.where(id: params[:id]).first
			if project
				{"data" => {"id" => project.id, "name" => project.name}, "status" => "Success"}
			else
				# error!({ "status" => "Fail", "error_message" => "Failed to save project" }, 404)
				{ "status" => "Fail", "error_message" => "Failed to save project" }
			end
		end
	end
end