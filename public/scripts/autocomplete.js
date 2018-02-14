google.maps.event.addDomListener(window, 'load', function () {
    var places = new google.maps.places.Autocomplete(document.getElementById('search'));
    google.maps.event.addListener(places, 'place_changed', function () {
      var place = places.getPlace();
      var address = place.formatted_address;
      var location = place.geometry.location;
      var lat = location.lat();
      var lon = location.lng();

      handleGeoData(lat, lon);
    });
});
