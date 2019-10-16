import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import MUIAppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

const styles = {
  root: {
    display: 'flex',
  },
  title: {
    flexGrow: 1,
  },
};

class TodoAppBar extends React.Component {

  // Open doc page
  openDoc = () => {
    var win = window.open("/swagger/index.html", '_blank');
    win.focus();
  };

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <MUIAppBar position="static" color="default">
          <Toolbar>
            <Typography variant="h6" className={classes.title}>Todo</Typography>
            <Button onClick={this.openDoc} color="inherit">TodoAPI Doc</Button>
          </Toolbar>
        </MUIAppBar>
      </div>
    )
  }
}

TodoAppBar.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(TodoAppBar);