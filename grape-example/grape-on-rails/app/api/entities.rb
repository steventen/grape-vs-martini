module GrapeApiRails
	module APIEntities
	  class Project < Grape::Entity
	    expose :id
	    expose :name
	  end
	end
end