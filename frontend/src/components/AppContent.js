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
        top: theme.spacing(4) * -1,
        right: theme.spacing(2),
    },
});

class AppContent extends React.Component {

    state = {
        openForm: false,
        showForm: false,
        editMode: false,
        todo: {}
    }

    componentDidMount() {
        this.props.dispatch(todoActions.get());
    }

    newForm = () => {
        this.setState({ openForm: true, showForm: true, editMode: false, todo: {} })
    }

    editForm = (todo) => {
        this.setState({ openForm: true, showForm: true, editMode: true, todo })
    }

    closeForm = () => {
        this.setState({ openForm: false, showForm: true })
        setTimeout(() => {
            this.setState({ showForm: false })
            this.props.dispatch(todoActions.get());
        }, 500);
    }

    render() {
        const { classes, data } = this.props;
        const { todo } = this.state;
        return (
            <main className={classes.layout}>
                <div className={classes.appBarSpacer} />
                <div className={classes.root}>

                    <Grid container spacing={1}>
                        {data && data.map((m, key) => (
                            <Grid item xs={12} sm={6} md={4} lg={3} key={key} onClick={() => this.editForm(m)} >
                                <AppCard todo={m} />
                            </Grid>)
                        )}
                        <Fab className={classes.fab} onClick={this.newForm} >
                            <AddIcon />
                        </Fab>
                    </Grid>
                    {this.state.showForm && <FormDialog open={this.state.openForm} edit={this.state.editMode} onClose={this.closeForm} todo={todo} />}
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