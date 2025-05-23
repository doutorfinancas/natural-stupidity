<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Vacation Calendar - Natural Stupidity</title>
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
  <link href='https://cdn.jsdelivr.net/npm/fullcalendar@6.1.8/main.min.css' rel='stylesheet' />
  <script src='https://cdn.jsdelivr.net/npm/fullcalendar@6.1.8/main.min.js'></script>
</head>
<body>
  <div class="container mt-5">
    <h1 class="mb-4">Vacation Calendar</h1>
    <div id='calendar'></div>
    <a href="/" class="btn btn-secondary mt-3">Dashboard</a>
  </div>
  <script>
  /* eslint-disable */
  document.addEventListener('DOMContentLoaded', function() {
    var events = [
    {{- range .Events }}
      { title: "{{ .Title }}", start: "{{ .Start }}", end: "{{ .End }}", color: "{{ .Color }}" },
    {{- end }}
    ];
    var calendarEl = document.getElementById('calendar');
    var calendar = new FullCalendar.Calendar(calendarEl, {
      initialView: 'dayGridMonth',
      events: events,
      height: 'auto'
    });
    calendar.render();
  });
  </script>
</body>
</html> 