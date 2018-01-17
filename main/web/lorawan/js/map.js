function show_map() {
  require([
    "esri/map",
    "esri/dijit/HomeButton",
    "dojo/domReady!"
  ], function (
    Map, HomeButton
  ) {
      map = new Map("map");
      // var layer = new esri.layers.ArcGISTiledMapServiceLayer("http://222.222.218.50:6080/arcgis/rest/services/MapServiceChina/MapServer");
      var layer = new esri.layers.ArcGISTiledMapServiceLayer("http://222.222.218.50:6080/arcgis/rest/services/MapServiceLuquanImg/MapServer");
      map.addLayer(layer)

      var home = new HomeButton({
        map: map
      }, "HomeButton");
      home.startup();

    });
}

function map_pan_left() {
  map.panLeft();
}

function map_pan_right() {
  map.panRight();
}

function map_pan_up() {
  map.panUp();
}

function map_pan_down() {
  map.panDown();
}

