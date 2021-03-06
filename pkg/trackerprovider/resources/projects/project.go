package projects

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
)

func NewProjectResource() *schema.Resource {
	return &schema.Resource{
		Create:        createProject,
		Read:          readProject,
		Delete:        deleteProject,
		Update:        updateProject,
		Exists:        existsProject,
		SchemaVersion: 1,
		Schema:        createSchema(),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func createProject(d *schema.ResourceData, meta interface{}) error {
	projectsRequest := pt.ProjectsRequest{}
	projectsRequest.AccountID = d.Get("account_id").(int)
	projectsRequest.AtomEnabled = d.Get("atom_enabled").(bool)
	projectsRequest.AutomaticPlanning = d.Get("automatic_planning").(bool)
	projectsRequest.BugsAndChoresAreEstimatable = d.Get("bugs_and_chores_are_estimatable").(bool)
	projectsRequest.Description = d.Get("description").(string)
	projectsRequest.EnableIncomingEmails = d.Get("enable_incoming_emails").(bool)
	projectsRequest.EnableTasks = d.Get("enable_tasks").(bool)
	projectsRequest.InitialVelocity = d.Get("initial_velocity").(int)
	projectsRequest.IterationLength = d.Get("iteration_length").(int)
	projectsRequest.JoinAs = d.Get("join_as").(string)
	projectsRequest.Name = d.Get("name").(string)
	projectsRequest.NumberOfDoneIterationsToShow = d.Get("number_of_done_iterations_to_show").(int)
	projectsRequest.PointScale = d.Get("point_scale").(string)
	projectsRequest.ProfileContent = d.Get("profile_content").(string)
	projectsRequest.ProjectType = d.Get("project_type").(string)
	projectsRequest.Public = d.Get("public").(bool)
	projectsRequest.Status = d.Get("status").(string)
	projectsRequest.VelocityAveragedOver = d.Get("velocity_averaged_over").(int)
	client := meta.(pt.ClientCaller)
	projectResponse, _, err := client.NewProject(projectsRequest)
	if err != nil {
		return fmt.Errorf("creating new project failed: %v", err)
	}

	d.SetId(strconv.Itoa(projectResponse.ID))
	return nil
}

func readProject(d *schema.ResourceData, meta interface{}) error {
	client := meta.(pt.ClientCaller)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("conversion of id failed: %v", err)
	}

	projectResponse, _, err := client.GetProject(id)
	if err != nil {
		return fmt.Errorf("get project api call failed: %v", err)
	}

	d.Set("account_id", projectResponse.AccountID)
	d.Set("atom_enabled", projectResponse.AtomEnabled)
	d.Set("automatic_planning", projectResponse.AutomaticPlanning)
	d.Set("bugs_and_chores_are_estimatable", projectResponse.BugsAndChoresAreEstimatable)
	d.Set("description", projectResponse.Description)
	d.Set("enable_incoming_emails", projectResponse.EnableIncomingEmails)
	d.Set("enable_tasks", projectResponse.EnableTasks)
	d.Set("initial_velocity", projectResponse.InitialVelocity)
	d.Set("iteration_length", projectResponse.IterationLength)
	d.Set("name", projectResponse.Name)
	d.Set("number_of_done_iterations_to_show", projectResponse.NumberOfDoneIterationsToShow)
	d.Set("point_scale", projectResponse.PointScale)
	d.Set("profile_content", projectResponse.ProfileContent)
	d.Set("project_type", projectResponse.ProjectType)
	d.Set("public", projectResponse.Public)
	d.Set("velocity_averaged_over", projectResponse.VelocityAveragedOver)
	d.SetId(strconv.Itoa(projectResponse.ID))
	return nil
}

func deleteProject(d *schema.ResourceData, meta interface{}) error {
	client := meta.(pt.ClientCaller)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("conversion of id failed: %v", err)
	}

	_, err = client.DeleteProject(id)
	if err != nil {
		return fmt.Errorf("delete project failed: %v", err)
	}

	return nil
}

func updateProject(d *schema.ResourceData, meta interface{}) error {
	projectRequest := pt.ProjectRequest{}
	projectRequest.AccountID = d.Get("account_id").(int)
	projectRequest.AtomEnabled = d.Get("atom_enabled").(bool)
	projectRequest.AutomaticPlanning = d.Get("automatic_planning").(bool)
	projectRequest.BugsAndChoresAreEstimatable = d.Get("bugs_and_chores_are_estimatable").(bool)
	projectRequest.Description = d.Get("description").(string)
	projectRequest.EnableIncomingEmails = d.Get("enable_incoming_emails").(bool)
	projectRequest.EnableTasks = d.Get("enable_tasks").(bool)
	projectRequest.InitialVelocity = d.Get("initial_velocity").(int)
	projectRequest.IterationLength = d.Get("iteration_length").(int)
	projectRequest.JoinAs = d.Get("join_as").(string)
	projectRequest.Name = d.Get("name").(string)
	projectRequest.NumberOfDoneIterationsToShow = d.Get("number_of_done_iterations_to_show").(int)
	projectRequest.PointScale = d.Get("point_scale").(string)
	projectRequest.ProfileContent = d.Get("profile_content").(string)
	projectRequest.ProjectType = d.Get("project_type").(string)
	projectRequest.Public = d.Get("public").(bool)
	projectRequest.Status = d.Get("status").(string)
	projectRequest.VelocityAveragedOver = d.Get("velocity_averaged_over").(int)
	client := meta.(pt.ClientCaller)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("conversion of id failed: %v", err)
	}

	projectResponse, _, err := client.UpdateProject(id, projectRequest)
	if err != nil {
		return fmt.Errorf("update project failed: %v", err)
	}

	d.SetId(strconv.Itoa(projectResponse.ID))
	return nil
}

func existsProject(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(pt.ClientCaller)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, fmt.Errorf("conversion of id failed: %v", err)
	}

	project, _, err := client.GetProject(id)
	if err != nil {
		return false, fmt.Errorf("get project api call failed: %v", err)
	}

	if project.ID > 0 {
		return true, nil
	}
	return false, nil
}

func createSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"no_owner": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Description: `
				boolean in the request body.
				 —  By default, the user whose credentials are supplied 
				 will be added as a project owner. To leave the project 
				 without this owner, supply the no_owner key.`,
		},

		"new_account_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `string[100] in the request body.
				 —  If specified, creates a new account with the specified 
				 name, and adds the new project to that account.`,
		},

		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			Description: `extended string[50] in the request body.
				 —  The name of the project.`,
		},

		"status": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				string in the request body.
				 —  The status of the project.`,
		},

		"iteration_length": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: `
				int in the request body.
				 —  The number of weeks in an iteration.`,
		},

		"week_start_day": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				enumerated string in the request body.
				 —  The day in the week the project's iterations are
				 to start on.	Valid enumeration values: Sunday, Monday, 
				 Tuesday, Wednesday, Thursday, Friday, Saturday`,
		},

		"point_scale": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: `
				string[255] in the request body.
				 —  The specification for the "point scale" available 
				 for entering story estimates within the project. It 
				 is specified as a comma-delimited series of values--any 
				 value that would be acceptable on the Project Settings 
				 page of the Tracker web application may be used here. 
				 If an exact match to one of the built-in point scales, 
				 the project will use that point scale. If another 
				 comma-separated point-scale string is passed, it will 
				 be treated as a "custom" point scale. The built-in 
				 scales are "0,1,2,3", "0,1,2,4,8", and "0,1,2,3,5,8".`,
		},

		"bugs_and_chores_are_estimatable": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Description: `
				boolean in the request body.
				 —  When true, Tracker will allow estimates to be set 
				 on Bug- and Chore-type stories. This is strongly not 
				 recommended. Please see the FAQ for more information.`,
		},

		"automatic_planning": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
			Description: `
				boolean in the request body.
				 —  When false, Tracker suspends the emergent planning of 
				 iterations based on the project's velocity, and allows 
				 users to manually control the set of unstarted stories 
				 included in the Current iteration. See the FAQ for more information.`,
		},

		"enable_tasks": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
			Description: `
				boolean in the request body.
				 —  When true, Tracker allows individual tasks to be 
				 created and managed within each story in the project.`,
		},

		"start_date": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				date in the request body.
				 —  The first day that should be in an iteration of the 
				 project. If both this and "week_start_day" are supplied,
				 they must be consistent. It is specified as a string in 
				 the format "YYYY-MM-DD" with "01" for January. If this is 
				 not supplied, it will remain blank (null), but "start_time"
				 will have a default value based on the stories in the project.
				 If a value is supplied for start_date, but that date is 
				 later than the accepted_at date of the earliest accepted 
				 story in your project, start_time will be based on the 
				 accepted_at date of the earliest accepted story.`,
		},

		"time_zone": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				time_zone in the request body.  —  The "native" time zone for the
				project, independent of the time zone(s) from which members of the
				project view or modify it.`,
		},

		"velocity_averaged_over": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: `
				int in the request body.  —  The number of iterations that should be used when
				averaging the number of points of Done stories in order to compute the
				project's velocity.`,
		},

		"number_of_done_iterations_to_show": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: `
				int in the request body.  —  There are areas within the Tracker UI and the API
				in which sets of stories automatically exclude the Done stories contained in
				older iterations. For example, in the web UI, the DONE panel doesn't
				necessarily show all Done stories by default, and provides a link to click to
				cause the full story set to be loaded/displayed. The value of this attribute is
				the maximum number of Done iterations that will be loaded/shown/included in
				these areas.`,
		},

		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				extended string[140] in the request body.  —  A description of the
				project's content. Entered through the web UI on the Project Settings
				page.`,
		},

		"profile_content": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				extended string[65535] in the request body.  —  A long description of
				the project. This is displayed on the Project Overview page in the
				Tracker web UI.`,
		},

		"enable_incoming_emails": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
			Description: `
				boolean in the request body.  —  When true, the project will accept
				incoming email responses to Tracker notification emails and convert
				them to comments on the appropriate stories.`,
		},

		"initial_velocity": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: `
				int in the request body.  —  The number which should be used as the
				project's velocity when there are not enough recent iterations with
				Done stories for an actual velocity to be computed.`,
		},

		"project_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: `
				enumerated string in the request body.  —  The project's type which
				determines visibility and permissions [demo is deprecated].  Valid
				enumeration values: demo, private, public, shared`,
		},

		"public": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Description: `
				boolean in the request body.  —  When true, Tracker will allow any user
				on the web to view the content of the project. The project will not
				count toward the limits of a paid subscription, and may be included on
				Tracker's Public Projects listing page.`,
		},

		"atom_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Description: `
				boolean in the request body.  —  When true, Tracker allows people to
				subscribe to the Atom (RSS, XML) feed of project changes.`,
		},

		"account_id": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: `
				int in the request body.  —  The ID number for the account which
				contains the project.`,
		},

		"join_as": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				enumerated string in the request body.  —  The default join_as value
				for the project [viewer, member].  Valid enumeration values: owner,
				member, viewer
 
				The new project is created with the currently-authenticated user as its
				original Owner. The server will supply a default value for any optional
				parameter that the request doesn't include. The default values are not defined
				as part of the API--they may change from time to time without an increase in
				the current API version number. Additionally, new project attributes may be
				added at any time without advance notice. The client will know what values the
				server has supplied from the response to the request. If the value set for a
				project attribute is important to an application, it should be included
				explicitly in the request without relying on the server to supply the
				default.`,
		},
	}
}
