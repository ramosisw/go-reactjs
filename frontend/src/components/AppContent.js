import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { todoActions } from '../redux/todo/todo.actions';
import FormDialog from './FormDialog';
import AppCard from './AppCard';
import Fab from '@material-ui/core/Fab';
import AddIcon from '@material-ui/icons/Add';
import { red } from '@material-ui/core/colors';

import { withStyles } from '@material-ui/core/styles';
import Grid from "@material-ui/core/Grid/Grid";

const styles = theme => ({
    root: {
        display: 'flex',
        position: 'relative',
        top: theme.spacing(2),
    },
    layout: {
        width: 'auto',
        marginLeft: theme.spacing(2),
        marginRight: theme.spacing(2),
        [theme.breakpoints.up(1080 + theme.spacing(2))]: {
            width: 1080,
            marginLeft: 'auto',
            marginRight: 'auto',
        },
    },
    appBarSpacer: theme.mixins.toolbar,
    fab: {
        backgroundColor: red[200],
        '&:hover': {
            backgroundColor: red[400],
        },
        position: 'absolute',
        bottom: theme.spacing(2),
        right: theme.spacing(2),
    },
});

class AppContent extends React.Component {

    state = {
        openForm: false
    }

    componentDidMount() {
        this.props.dispatch(todoActions.get());
    }

    openForm = () => {
        this.setState({ openForm: true })
    }

    openDetail = () => {
        this.setState({ openForm: true })
    }

    closeForm = () => {
        this.setState({ openForm: false })
    }

    render() {
        const { classes, data } = this.props;
        return (
            <main className={classes.layout}>
                <div className={classes.appBarSpacer} />
                <div className={classes.root}>

                    <Grid container spacing={1}>
                        {data && data.map((m, key) => (
                            <Grid item md={4} key={key} onClick={this.openDetail} > 
                                <AppCard todo={m} />
                            </Grid>)
                        )}
                        <Fab className={classes.fab} onClick={this.openForm} >
                            <AddIcon />
                        </Fab>
                    </Grid>
                    <FormDialog open={this.state.openForm} onClose={this.closeForm} />
                </div>
            </main>
        )
    }
}

function mapStateToProps(state) {
    const { data } = state.todo.get;
    return {
        data
    };
}

AppContent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(connect(mapStateToProps)(AppContent));