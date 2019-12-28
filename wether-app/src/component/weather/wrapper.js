import React from "react";
import "weather-icons/css/weather-icons.css";

import Form from "./form.component";
import Weather from "./weather.component";
import {fetchCurrentWeather} from "../../api/request";


const Api_Key = "429736441cf3572838aa10530929f7cd";

class Wrapper extends React.Component {
  constructor() {
    super();
    this.state = {
      city: undefined,
      country: undefined,
      icon: undefined,
      main: undefined,
      celsius: undefined,
      temp_max: null,
      temp_min: null,
      description: "",
      error: false,
      error_msg : ""
    };

    this.weatherIcon = {
      Thunderstorm: "wi-thunderstorm",
      Drizzle: "wi-sleet",
      Rain: "wi-storm-showers",
      Snow: "wi-snow",
      Atmosphere: "wi-fog",
      Clear: "wi-day-sunny",
      Clouds: "wi-day-fog"
    };
  }

  get_WeatherIcon(icons, rangeId) {
    switch (true) {
      case rangeId >= 200 && rangeId < 232:
        this.setState({ icon: icons.Thunderstorm });
        break;
      case rangeId >= 300 && rangeId <= 321:
        this.setState({ icon: icons.Drizzle });
        break;
      case rangeId >= 500 && rangeId <= 521:
        this.setState({ icon: icons.Rain });
        break;
      case rangeId >= 600 && rangeId <= 622:
        this.setState({ icon: icons.Snow });
        break;
      case rangeId >= 701 && rangeId <= 781:
        this.setState({ icon: icons.Atmosphere });
        break;
      case rangeId === 800:
        this.setState({ icon: icons.Clear });
        break;
      case rangeId >= 801 && rangeId <= 804:
        this.setState({ icon: icons.Clouds });
        break;
      default:
        this.setState({ icon: icons.Clouds });
    }
  }

  calCelsius(temp) {
    let cell = Math.floor(temp - 273.15);
    return cell;
  }

  getWeather = async e => {
    e.preventDefault();

    const country = e.target.elements.country.value;
    const city = e.target.elements.city.value;

    if (country && city) {
        try {
            const token = await fetchCurrentWeather({country: country, city: city}, this.state.token);

            this.props.onAuthenticated(token);
            this.props.history.push("/");
        } catch (e) {
            console.log(e);
        }

      const response = await api_call.json();

      if(response.cod !== "200") {
        this.setState({
          city: undefined,
          country: undefined,
          icon: undefined,
          main: undefined,
          celsius: undefined,
          temp_max: null,
          temp_min: null,
          description: "",
          error: true,
          error_msg: response.message 
        });
        return;
      }

      this.get_WeatherIcon(this.weatherIcon, response.weather[0].id);

      this.setState({
        city: `${response.name}, ${response.sys.country}`,
        country: response.sys.country,
        main: response.weather[0].main,
        celsius: this.calCelsius(response.main.temp),
        temp_max: this.calCelsius(response.main.temp_max),
        temp_min: this.calCelsius(response.main.temp_min),
        description: response.weather[0].description,
        error: false
      });

      console.log(response);
    } else {
      this.setState({
        city: undefined,
          country: undefined,
          icon: undefined,
          main: undefined,
          celsius: undefined,
          temp_max: null,
          temp_min: null,
          description: "",
          error: true,
          error_msg: "Please Enter City and Country...!"
      });
    }
  };

  render() {
    if (this.props.token) {
    return (<div>
            <Form loadweather={this.getWeather} error={this.state.error} error_msg={this.state.error_msg} />
            <Weather
              cityname={this.state.city}
              weatherIcon={this.state.icon}
              temp_celsius={this.state.celsius}
              temp_max={this.state.temp_max}
              temp_min={this.state.temp_min}
              description={this.state.description}
            />   
            </div>     
    );
    } 
    return (
        <Redirect to="/login" />
    );
  }
}

export default Wrapper;

