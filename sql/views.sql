DROP VIEW IF EXISTS vw_artifacts;
DROP VIEW IF EXISTS vw_tasks;
DROP VIEW IF EXISTS vw_builds;

CREATE VIEW vw_builds AS
SELECT
   repo.name as repo_name,
   build.id as build_id, 
   "build"."created" AS "build_created", "build"."duration" AS "build_duration", 
   "build"."clock_duration" AS "build_clock_duration", 
   "build"."branch" AS "build_branch", "build"."pull_request" AS "build_pull_request",
   build.status AS build_status
FROM  builds as build
LEFT JOIN repos as repo ON (build.repo_id = repo.id);



CREATE VIEW vw_tasks  AS
 SELECT 
   build.*,
   "task"."id" as task_id,
   "task"."name" as task_name, "task"."duration" as task_duration, "task"."created" as task_created, "task"."scheduled" as task_schedule, "task"."executing" as task_executing, task.status as task_status
  FROM "tasks" AS "task" 
   LEFT JOIN "vw_builds" AS "build" ON ("build"."build_id" = "task"."build_id")
;


create VIEW vw_artifacts AS
  SELECT 
  task.*, 
  artifact.id as artifact_id, 
  artifact.location as artifact_location
  FROM artifacts as artifact
  LEFT JOIN vw_tasks as task on (task.task_id = artifact.task_id);

